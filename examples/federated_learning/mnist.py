"""Adapted from the PyTorch Lightning quickstart example.

Source: https://pytorchlightning.ai/ (2021/02/04)
"""

import numpy as np
from collections import OrderedDict

import torch
from torch import nn
from torch.nn import functional as F
import pytorch_lightning as pl

class LitAutoEncoder(pl.LightningModule):
    def __init__(self):
        super().__init__()
        self.encoder = nn.Sequential(
            nn.Linear(28 * 28, 64),
            nn.ReLU(),
            nn.Linear(64, 3),
        )
        self.decoder = nn.Sequential(
            nn.Linear(3, 64),
            nn.ReLU(),
            nn.Linear(64, 28 * 28),
        )

    def forward(self, x):
        embedding = self.encoder(x)
        return embedding

    def configure_optimizers(self):
        optimizer = torch.optim.Adam(self.parameters(), lr=1e-3)
        return optimizer

    def training_step(self, train_batch, batch_idx):
        x, y = train_batch
        x = x.view(x.size(0), -1)
        z = self.encoder(x)
        x_hat = self.decoder(z)
        loss = F.mse_loss(x_hat, x)
        self.log("train_loss", loss)
        return loss

    def validation_step(self, batch, batch_idx):
        self._evaluate(batch, "val")

    def test_step(self, batch, batch_idx):
        self._evaluate(batch, "test")

    def _evaluate(self, batch, stage=None):
        x, y = batch
        x = x.view(x.size(0), -1)
        z = self.encoder(x)
        x_hat = self.decoder(z)
        loss = F.mse_loss(x_hat, x)
        if stage:
            self.log(f"{stage}_loss", loss, prog_bar=True)

def load_parameters(parameters_path):
    try:
        parameters = np.load(parameters_path, allow_pickle=True).tolist()
        return parameters
    except OSError:
        return None

def save_parameters(parameters_path, parameters):
    np.save(parameters_path, np.array(parameters, dtype=object))

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
