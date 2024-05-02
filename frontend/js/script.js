document.addEventListener('DOMContentLoaded', () => {
  initializeGameBoard();
});

function initializeGameBoard() {
  const board = document.getElementById('game-board');
  for (let i = 0; i < 19; i++) { // Assuming a standard Catan board of 19 hex tiles
      const tile = document.createElement('div');
      tile.className = 'tile';
      tile.innerHTML = `<span class="resource">${getResourceType(i)}</span>`;
      board.appendChild(tile);
  }
}

function getResourceType(i) {
  const resources = ['Wood', 'Brick', 'Sheep', 'Wheat', 'Ore', 'Desert'];
  return resources[i % resources.length];
}

document.getElementById('roll-dice').addEventListener('click', () => {
  const roll = Math.floor(Math.random() * 6) + 1 + Math.floor(Math.random() * 6) + 1;
  document.getElementById('resource-list').innerText = `Dice Roll: ${roll}`;
  // Further logic to distribute resources based on the roll
});

document.getElementById('end-turn').addEventListener('click', () => {
  alert('Turn ended. Next player’s turn.');
  // Implement turn change logic here
});
