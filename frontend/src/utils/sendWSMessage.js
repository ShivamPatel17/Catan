export function SendWSMessage(socket, message) {
	if (socket && socket.readyState === WebSocket.OPEN) {
		socket.send(JSON.stringify(message));
	} else {
		console.error("WebSocket connection is not open", "failed to send this data: ", message);
	}
}
