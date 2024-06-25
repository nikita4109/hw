// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract UserWallet {
    address public owner;
    uint256 public balance;
    mapping(string => uint256) public assets;

    constructor(address _owner) {
        owner = _owner;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }

    function deposit() public payable onlyOwner {
        balance += msg.value;
    }

    function buyAsset(string memory asset, uint256 amount) public onlyOwner {
        require(balance >= amount, "Insufficient balance");
        balance -= amount;
        assets[asset] += amount;
    }

    function sellAsset(string memory asset, uint256 amount) public onlyOwner {
        require(assets[asset] >= amount, "Insufficient asset balance");
        assets[asset] -= amount;
        balance += amount;
    }

    function getBalance() public view onlyOwner returns (uint256) {
        return balance;
    }

    function getAssetBalance(string memory asset) public view onlyOwner returns (uint256) {
        return assets[asset];
    }
}

contract WalletFactory {
    mapping(address => address) public userWallets;

    function createWallet() public {
        require(userWallets[msg.sender] == address(0), "Wallet already exists");
        UserWallet wallet = new UserWallet(msg.sender);
        userWallets[msg.sender] = address(wallet);
    }

    function getWallet(address user) public view returns (address) {
        return userWallets[user];
    }
}
