# ASM-env
This project is part of [asm-webhook](https://github.com/ayoul3/asm-webhook). It is the binary that gets injected into containers, decrypts SecretsManager environment variables and spawns the target command with a fresh batch of decrypted secrets.

## Credit
[Banzai Vaults](https://github.com/banzaicloud/bank-vaults/tree/master/charts/vault-secrets-webhook)