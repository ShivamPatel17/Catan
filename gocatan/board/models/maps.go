package models

func RegularBoard() RelativeHexagonTile {
	startingHexX, startingHexY := 500.0, 450.0
	// can use this to pass configuation from the client in the future
	return RelativeHexagonTile{
		AdjacentTiles: []*DirectionalHexagonTile{
			{
				Direction: TopRight,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       MiddleRight,
							RelativeHexTile: RelativeHexagonTile{},
						},
						{
							Direction:       TopRight,
							RelativeHexTile: RelativeHexagonTile{},
						},
						{
							Direction:       TopLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				Direction: TopLeft,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       TopLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
						{
							Direction:       MiddleLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				Direction: MiddleRight,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       MiddleRight,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				Direction: MiddleLeft,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       MiddleLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				Direction: BottomRight,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       MiddleRight,
							RelativeHexTile: RelativeHexagonTile{},
						},
						{
							Direction: BottomRight,
							RelativeHexTile: RelativeHexagonTile{
								AdjacentTiles: []*DirectionalHexagonTile{
									{
										Direction:       MiddleRight,
										RelativeHexTile: RelativeHexagonTile{},
									},
								},
							},
						},
						{
							Direction:       BottomLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				Direction: BottomLeft,
				RelativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []*DirectionalHexagonTile{
						{
							Direction:       MiddleLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
						{
							Direction:       BottomLeft,
							RelativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
		},
		Concrete: &ConcreteHexagonTile{
			X:        startingHexX,
			Y:        startingHexY,
			Resource: Sheep,
		},
	}
}

func SingleHex() RelativeHexagonTile {
	startingHexX, startingHexY := 500.0, 450.0
	// can use this to pass configuation from the client in the future
	return RelativeHexagonTile{
		Concrete: &ConcreteHexagonTile{
			X:        startingHexX,
			Y:        startingHexY,
			Resource: Sheep,
		},
	}

}

// can use this to pass configuation from the client in the future
