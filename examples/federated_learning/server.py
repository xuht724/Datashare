import argparse
from logging import error
import numpy as np
import os

def list_all_parameters(path):    
    parameters_list = []
    if os.path.exists(path):
        files=os.listdir(path)
    else:
        print('this path not exist')
    for file in files:
        if os.path.isdir(os.path.join(path,file)):
            parameters_list += list_all_parameters(os.path.join(path,file))
        else:
            if file.find("parameters.npy") != -1:
                parameters = np.load(os.path.join(path,file), allow_pickle=True).tolist()
                parameters_list.append(parameters)
    return parameters_list

def fedavg(parameters_list):
    if len(parameters_list) <= 0:
        error("no client parameters!!!")
    list = np.array(parameters_list, dtype=object)
    sum = np.sum(list,0)
    avg = np.divide(sum, len(list))
    return avg.tolist()

def main() -> None:
    parser = argparse.ArgumentParser(description='PyTorch MNIST Example -- Server')
    parser.add_argument('--path', type=str, default="results/", metavar='N',
                        help='the path of clients\' parameters file(default: "results/")')
    args = parser.parse_args()

    # create directory
    dir = 'results/avg/'
    if not os.path.exists(dir):
        os.makedirs(dir)

    # get parameters
    parameters_list = list_all_parameters(args.path)

    # aggregate 
    avg_parameters = fedavg(parameters_list)

    # save
    np.save(dir+'parameters.npy', np.array(avg_parameters, dtype=object))

if __name__ == "__main__":
    main()
