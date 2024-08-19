package board

func RegularBoard() RelativeHexagonTile {
	startingHexX, startingHexY := 500, 450
	// can use this to pass configuation from the client in the future
	return RelativeHexagonTile{
		AdjacentTiles: []DirectionalHexagonTile{
			{
				direction: TopRight,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       TopRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       TopLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: TopLeft,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       TopLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: MiddleRight,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: MiddleLeft,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: BottomRight,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: BottomLeft,
				relativeHexTile: RelativeHexagonTile{
					AdjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomLeft,
							relativeHexTile: RelativeHexagonTile{},
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

// can use this to pass configuation from the client in the future
