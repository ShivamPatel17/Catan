var config = {
  type: Phaser.AUTO,
  width: 800,
  height: 600,
  physics: {
    default: 'arcade',
    arcade: {
        debug: true
    }
},
  scene: {
      preload: preload,
      create: create,
      update: update
  }
};

var game = new Phaser.Game(config);
var cursors;
var dice;
var die;

var dieDim = 68


function preload() { 

  this.load.spritesheet('redDie', 'boardgamePack_v2/Spritesheets/diceRed_border.png', {frameWidth: dieDim, frameHeight: dieDim})
}

function create(){
  dice = this.physics.add.staticGroup();
  die = dice.create(300, 300, 'redDie')


  cursors = this.input.keyboard.createCursorKeys();
}

function update(){
  if (cursors.left.isDown) { 
    roll()
  }
}

function roll() { 
  die.disableBody(true, true)
}