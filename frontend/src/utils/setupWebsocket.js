export function SetupWebSocket(scene) {
	// Create a new WebSocket connection
	scene.socket = new WebSocket("ws://localhost:3000/ws");

	// Handle the connection opening
	scene.socket.onopen = () => {
		console.log("WebSocket connection established");
	};

	// Handle incoming messages from the server
	scene.socket.onmessage = (event) => {
		// Check if the message is a Blob
		if (event.data instanceof Blob) {
			event.data.text().then((text) => {
				try {
					const message = JSON.parse(text); // Parse the string as JSON
					console.log("Parsed JSON message:", message); // Log the parsed JSON
					scene.handleServerMessage(message); // Call your handler with the parsed message
				} catch (error) {
					console.error("Error parsing JSON:", error);
				}
			});
		} else {
			// If the message is not a Blob, handle it as a string (if needed)
			try {
				const message = JSON.parse(event.data); // Parse it directly as JSON
				console.log("Parsed JSON message:", message);
				scene.handleServerMessage(message);
			} catch (error) {
				console.error("Error parsing JSON:", error);
			}
		}
	};

	// Handle connection closure
	scene.socket.onclose = () => {
		console.log("WebSocket connection closed");
	};

	// Handle connection errors
	scene.socket.onerror = (error) => {
		console.error("WebSocket error:", error);
	};
}
