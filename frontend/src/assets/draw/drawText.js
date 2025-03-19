export function DrawText(scene, y, x, text) {
	scene.add.text(y, x, text, {
		fontSize: '16px',
		fill: '#fff',
		fontFamily: 'Arial'
	});
}
