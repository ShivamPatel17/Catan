import { DrawText } from "../../assets/draw/drawText"

export function CreateUserData(scene) {
	DrawText(scene, 800, 100, scene.playerUuid)
}
