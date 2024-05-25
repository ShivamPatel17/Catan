// the game itself
let game;

// global object with game options
let gameOptions = {
  // card width, in pixels
  diceWidth: 64,
  // card height, in pixels
  diceHeight: 64,
  // card scale. 1 = original size, 0.5 half size and so on
  cardScale: 0.8,
};

window.onload = function () {
  let gameConfig = {
    type: Phaser.AUTO,
    backgroundColor: 0x4488aa,
    scale: {
      mode: Phaser.Scale.FIT,
      autoCenter: Phaser.Scale.CENTER_BOTH,
      parent: "thegame",
      width: 800,
      height: 600,
    },
    scene: playGame,
  };
  game = new Phaser.Game(gameConfig);
  window.focus();
};

// Function to make a GET request
async function fetchData(url) {
  try {
    const response = await fetch(url);
    // Checking if the response is ok (status 200-299)
    if (!response.ok) {
      throw new Error("Network response was not ok " + response.statusText);
    }
    const data = await response.json(); // Parsing JSON data from the response
    return data;
  } catch (error) {
    console.error("There was a problem with the fetch operation: ", error);
  }
}

class playGame extends Phaser.Scene {
  constructor() {
    super("PlayGame");
    this.redDieNum = 1;
  }

  preload() {
    // loading the sprite sheet with all cards
    this.load.spritesheet(
      "redDie",
      "boardgamePack_v2/Spritesheets/diceRed.png",
      {
        frameWidth: gameOptions.diceWidth,
        frameHeight: gameOptions.diceHeight,
      }
    );
    // load hexagons
    this.load.image("hexagon", "assets/images/hexagon.png");
  }

  create() {
    this.die = this.add.sprite(700, 550, "redDie").setInteractive();
    this.input.keyboard.on("keydown-SPACE", this.rollDie, this);
    this.loadhex();
  }

  // Asynchronous function to roll the die
  async rollDie() {
    try {
      // Fetch the random number from the backend
      const number = await fetchData("http://localhost:3000/roll");

      console.log("Random number from backend:", number);

      // Set the dice sprite to display the randomly selected face
      this.die.setFrame(number - 1); // Subtracting 1 because frame indexing starts at 0
    } catch (error) {
      console.error("Error rolling die:", error);
    }
  }

  loadhex() {
    // Create multiple sprites in a loop
    for (let i = 0; i < 10; i++) {
      // Calculate x and y positions for each sprite
      let x = 100 + i * 60;
      let y = 300;

      // Create a sprite and add it to the scene
      const sprite = this.add.sprite(x, y, "hexagon");

      // Customize the sprite if needed
      sprite.setScale(1.0);

      // Store the sprite in the array for later use
      this.sprites.push(sprite);
    }
  }
}
