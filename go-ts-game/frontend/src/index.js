import * as Phaser from "phaser";
import GameScene from './gameScene';
var config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    scene: GameScene, // 使用默认导入的 GameScene 类
};
var game = new Phaser.Game(config);
