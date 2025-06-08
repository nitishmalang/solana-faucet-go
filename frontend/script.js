const connectWalletBtn = document.getElementById('connectWallet');
const requestSolBtn = document.getElementById('requestSol');
const statusText = document.getElementById('status');
let walletPublicKey = null;

connectWalletBtn.addEventListener('click', async () => {
    try {
        const provider = window.solana;
        if (!provider) {
            statusText.textContent = 'Please install Phantom Wallet!';
            return;
        }
        const resp = await provider.connect();
        walletPublicKey = resp.publicKey.toString();
        statusText.textContent = `Connected: ${walletPublicKey.slice(0, 10)}...`;
        connectWalletBtn.disabled = true;
        requestSolBtn.disabled = false;
    } catch (err) {
        statusText.textContent = 'Connection failed: ' + err.message;
    }
});

requestSolBtn.addEventListener('click', async () => {
    if (!walletPublicKey) {
        statusText.textContent = 'Please connect your wallet first!';
        return;
    }
    try {
        statusText.textContent = 'Requesting 1 SOL...';
        const response = await fetch('http://localhost:8082/airdrop', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ wallet_address: walletPublicKey })
        });
        const result = await response.json();
        statusText.textContent = result.message;
        if (result.status === 'success') {
            requestSolBtn.disabled = true;
        }
    } catch (err) {
        statusText.textContent = 'Airdrop failed: ' + err.message;
    }
});