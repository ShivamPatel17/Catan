// the game itself
let game;

// global object with game options
let gameOptions = {

    // card width, in pixels
    diceWidth: 64,

    // card height, in pixels
    diceHeight: 64,

    // card scale. 1 = original size, 0.5 half size and so on
    cardScale: 0.8
}
window.onload = function() {
    let gameConfig = {
        type: Phaser.AUTO,
        backgroundColor: 0x4488aa,
        scale: {
            mode: Phaser.Scale.FIT,
            autoCenter: Phaser.Scale.CENTER_BOTH,
            parent: "thegame",
            width: 800,
            height: 600
        },
        scene: playGame
    }
    game = new Phaser.Game(gameConfig);
    window.focus();
}

class playGame extends Phaser.Scene {
    constructor() {
        super("PlayGame");
        this.redDieNum = 1
    }

    preload() {
      // loading the sprite sheet with all cards
      this.load.spritesheet("redDie", "boardgamePack_v2/Spritesheets/diceRed.png", {
          frameWidth: gameOptions.diceWidth,
          frameHeight: gameOptions.diceHeight
      });
    }
    
    create() {
      // Add the dice sprite to the scene
      this.die = this.add.sprite(400, 300, 'redDie').setInteractive();

      // Add a keyboard listener for the spacebar
      this.input.keyboard.on('keydown-SPACE', this.rollDie, this);
  }

  rollDie() {
      // Generate a random dice face between 0 and 5 (for a six-faced dice)
      const randomFace = Phaser.Math.Between(0, 5);
      // Set the dice sprite to display the randomly selected face
      this.die.setFrame(randomFace);
  }

}

