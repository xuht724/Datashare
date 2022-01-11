"""Adapted from the PyTorch Lightning quickstart example.

Source: https://pytorchlightning.ai/ (2021/02/04)
"""


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

def main() -> None:
    """Centralized training."""

    # Load data
    train_loader, val_loader, test_loader = load_data()

    # Load model
    model = LitAutoEncoder()

    # Train
    trainer = pl.Trainer(max_epochs=5, progress_bar_refresh_rate=0)
    trainer.fit(model, train_loader, val_loader)

    # Test
    trainer.test(model, test_loader)


if __name__ == "__main__":
    main()
