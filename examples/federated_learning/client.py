import argparse
import numpy as np
import os
import mnist
import pytorch_lightning as pl
from collections import OrderedDict

import torch
from torch.utils.data import DataLoader, random_split, Subset
from torchvision import transforms
from torchvision.datasets import MNIST

def load_data(id):
    def id2range(id):
        if(id == 0):
            return range(0,5)
        elif(id == 1):
            return range(5,10)
    id_range = id2range(id)
    
    # Training / validation set
    trainset = MNIST("", train=True, download=True, transform=transforms.ToTensor())
    index = [i for i,cl in enumerate(trainset.targets.tolist()) if cl in id_range]
    trainset = Subset(trainset,index)

    length = len(index)
    a=int(length*0.9)
    b=length-a
    mnist_train, mnist_val = random_split(trainset, [a, b])
    train_loader = DataLoader(mnist_train, batch_size=32, shuffle=True, num_workers=4)
    val_loader = DataLoader(mnist_val, batch_size=32, shuffle=False, num_workers=4)

    # Test set
    testset = MNIST("", train=False, download=True, transform=transforms.ToTensor())
    testset = Subset(testset,range(1000))
    test_loader = DataLoader(testset, batch_size=32, shuffle=False, num_workers=4)

    return train_loader, val_loader, test_loader

def load_parameters(parameters_path):
    try:
        parameters = np.load(parameters_path, allow_pickle=True).tolist()
        return parameters
    except OSError:
        return None

def get_parameters(model):
    def _get_parameters(m):
        return [val.cpu().numpy() for _, val in m.state_dict().items()]
    encoder_params = _get_parameters(model.encoder)
    decoder_params = _get_parameters(model.decoder)
    return encoder_params + decoder_params

def set_parameters(model, parameters):
    def _set_parameters(m, p):
        params_dict = zip(m.state_dict().keys(), p)
        state_dict = OrderedDict({k: torch.tensor(v) for k, v in params_dict})
        m.load_state_dict(state_dict, strict=True)
        
    if parameters != None:
        _set_parameters(model.encoder, parameters[:4])
        _set_parameters(model.decoder, parameters[4:])

def main() -> None:
    parser = argparse.ArgumentParser(description='PyTorch MNIST Example -- Client')
    parser.add_argument('--clientID', type=int, default=0, metavar='N',
                        help='input clientID (default: 0)')
    parser.add_argument('--path', type=str, default="results/avg/parameters.npy", metavar='N',
                        help='the path of initial parameters file(default: "results/avg/parameters.npy")')
    args = parser.parse_args()
    
    # create directory
    dir = 'results/client-{}/'.format(args.clientID)
    if not os.path.exists(dir):
        os.makedirs(dir)

    # model and data
    model = mnist.LitAutoEncoder()
    train_loader, val_loader, test_loader = load_data(args.clientID)

    # initial parameters
    parameters = load_parameters(args.path)
    set_parameters(model, parameters)

    # train
    trainer = pl.Trainer(max_epochs=1, max_steps=10, enable_progress_bar=True)
    trainer.fit(model, train_loader, val_loader)

    # test
    results = trainer.test(model, test_loader)
    loss = results[0]["test_loss"]
    with open(dir+'loss',"a") as f:
        np.savetxt(f, [loss])

    # save
    parameters = get_parameters(model)
    np.save(dir+'parameters.npy', np.array(parameters, dtype=object))
  
if __name__ == "__main__":
    main()