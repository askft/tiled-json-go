package tiled

import (
	"encoding/json"
	"io/ioutil"
)

/*
	Based on the JSON Map Format for Tiled 1.2.
	https://doc.mapeditor.org/en/stable/reference/json-map-format/

	TODO:
		- This needs to be extensively tested. You can help by e.g.
		  sending me your Tiled files (as JSON) in order to give me
		  testing material.
		- Comment all fields with the description in the link above.
		  Textually separate optional fields and describe them as such.
		- Add the omitempty tag where appropriate.
*/

// ParseTilemap takes a Tiled tilemap JSON file
// and converts it into a Golang data structure.
func ParseTilemap(path string) (*Map, error) {
	m := &Map{}
	jdata, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jdata, m)
	return m, err
}

// ParseTileset takes a Tiled tileset JSON file
// and converts it into a Golang data structure.
func ParseTileset(path string) (*Tileset, error) {
	ts := &Tileset{}
	jdata, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jdata, ts)
	return ts, err
}

// Map describes a Tiled map.
type Map struct {
	BackgroundColor string     `json:"backgroundcolor"` // Hex-formatted color (#RRGGBB or #AARRGGBB) (optional).
	Height          int        `json:"height"`          // Number of tile rows.
	HexSideLength   int        `json:"hexsidelength"`   // Length of the side of a hex tile in pixels.
	Infinite        bool       `json:"infinite"`        // Whether the map has infinite dimensions.
	Layers          []Layer    `json:"layers"`          // Array of Layers.
	NextLayerID     int        `json:"nextlayerid"`     // Auto-increments for each layer.
	NextObjectID    int        `json:"nextobjectid"`    // Auto-increments for each placed object.
	Orientation     string     `json:"orientation"`     // "orthogonal", "isometric", "staggered" or "hexagonal".
	Properties      []Property `json:"properties"`      // A list of properties (name, value, type).
	RenderOrder     string     `json:"renderorder"`     // Rendering direction (orthogonal maps only).
	StaggerAxis     string     `json:"staggeraxis"`     // "x" or "y" (staggered / hexagonal maps only).
	StaggerIndex    string     `json:"staggerindex"`    // "odd" or "even" (staggered / hexagonal maps only).
	TiledVersion    string     `json:"tiledversion"`    // The Tiled version used to save the file.
	TileHeight      int        `json:"tileheight"`      // Map grid height.
	Tilesets        []Tileset  `json:"tilesets"`        // Array of Tilesets.
	TileWidth       int        `json:"tilewidth"`       // Map grid width.
	Type            string     `json:"type"`            // "map" (since 1.0).
	Version         float64    `json:"version"`         // The JSON format version.
	Width           int        `json:"width"`           // Number of tile columns.
}

// Property TODO describe
type Property struct {
	Name  string      `json:"name"`  // Name of property
	Type  string      `json:"type"`  // Type of property value
	Value interface{} `json:"value"` // Value of property
}

// Layer TODO describe
type Layer struct {
	// Common
	ID         int        `json:"id"`         // Incremental id - unique across all layers
	Name       string     `json:"name"`       // Name assigned to this layer
	Type       string     `json:"type"`       // "tilelayer, "objectgroup, "imagelayer or "group"
	Visible    bool       `json:"visible"`    // Whether layer is shown or hidden in editor
	Width      int        `json:"width"`      // Column count. Same as map width for fixed-size maps
	Height     int        `json:"height"`     // Row count. Same as map height for fixed-size maps
	X          int        `json:"x"`          // Horizontal layer offset in tiles. Always 0
	Y          int        `json:"y"`          // Vertical layer offset in tiles. Always 0
	OffsetX    float64    `json:"offsetx"`    // Horizontal layer offset in pixels (default: 0)
	OffsetY    float64    `json:"offsety"`    // Vertical layer offset in pixels (default: 0)
	Opacity    float64    `json:"opacity"`    // Value between 0 and 1
	Properties []Property `json:"properties"` // A list of properties (name, value, type)

	// TileLayer only
	Chunks      []Chunk     `json:"chunks"`      // Array of chunks (optional, for ininite maps)
	Compression string      `json:"compression"` // "zlib", "gzip" or empty (default)
	Data        interface{} `json:"data"`        // Array or string. Array of unsigned int (GIDs) or base64-encoded data
	Encoding    string      `json:"encoding"`    // "csv" (default) or "base64"

	// ObjectGroup only
	Objects   []Object `json:"objects"`   // Array of objects
	DrawOrder string   `json:"draworder"` // "topdown" (default) or "index"

	// Group only
	Layers []Layer `json:"layers"` // Array of layers

	// ImageLayer only
	Image            string `json:"image"`            // Image used by this layer
	TransparentColor string `json:"transparentcolor"` // Hex-formatted color (#RRGGBB) (optional)
}

// Chunk is used to store the tile layer data for infinite maps.
type Chunk struct {
	Data   interface{} `json:"data"`   // Array of unsigned int (GIDs) or base64-encoded data
	Height int         `json:"height"` // Height in tiles
	Width  int         `json:"width"`  // Width in tiles
	X      int         `json:"x"`      // X coordinate in tiles
	Y      int         `json:"y"`      // Y coordinate in tiles
}

// Object TODO describe
type Object struct {
	ID         int                    `json:"id"`                 // Incremental id - unique across all objects
	GID        int                    `json:"gid"`                // GID, only if object comes from a Tilemap
	Name       string                 `json:"name"`               // String assigned to name field in editor
	Type       string                 `json:"type"`               // String assigned to type field in editor
	X          float64                `json:"x"`                  // X coordinate in pixels
	Y          float64                `json:"y"`                  // Y coordinate in pixels
	Width      float64                `json:"width"`              // Width in pixels, ignored if using a gid
	Height     float64                `json:"height"`             // Height in pixels, ignored if using a gid
	Visible    bool                   `json:"visible"`            // Whether object is shown in editor
	Ellipse    bool                   `json:"ellipse"`            // Used to mark an object as an ellipse
	Point      bool                   `json:"point"`              // Used to mark an object as a point
	Polygon    []Coordinate           `json:"polygon"`            // A list of x,y coordinates in pixels
	Polyline   []Coordinate           `json:"polyline"`           // A list of x,y coordinates in pixels
	Properties []Property             `json:"properties"`         // A list of properties (name, value, type)
	Rotation   float64                `json:"rotation"`           // Angle in degrees clockwise
	Template   string                 `json:"template,omitempty"` // Reference to a template file, in case object is a template instance
	Text       map[string]interface{} `json:"text"`               // String key-value pairs
}

// TODO, read up about templates here http://docs.mapeditor.org/en/stable/manual/using-templates/

// Coordinate TODO describe
type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Offset TODO describe
type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// A Tileset that associates information with each tile, like its image
// path or terrain type, may include a Tiles array property. Each tile
// has an ID property, which specifies the local ID within the tileset.
//
// For the terrain information, each value is a length-4 array where
// each element is the index of a terrain on one corner of the tile.
// The order of indices is: top-left, top-right, bottom-left, bottom-right.
type Tileset struct {
	FirstGID         int        `json:"firstgid"`         // GID corresponding to the first tile in the set
	Source           string     `json:"source"`           // Only used if an external tileset is referred to
	Name             string     `json:"name"`             // Name given to this tileset
	Type             string     `json:"type"`             // "tileset" (for tileset files, since 1.0)
	Columns          int        `json:"columns"`          // The number of tile columns in the tileset
	Image            string     `json:"image"`            // Image used for tiles in this set
	ImageWidth       int        `json:"imagewidth"`       // Width of source image in pixels
	ImageHeight      int        `json:"imageheight"`      // Height of source image in pixels
	Margin           int        `json:"margin"`           // Buffer between image edge and first tile (pixels)
	Spacing          int        `json:"spacing"`          // Spacing between adjacent tiles in image (pixels)
	TileCount        int        `json:"tilecount"`        // The number of tiles in this tileset
	TileWidth        int        `json:"tilewidth"`        // Maximum width of tiles in this set
	TileHeight       int        `json:"tileheight"`       // Maximum height of tiles in this set
	TransparentColor string     `json:"transparentcolor"` // Hex-formatted color (#RRGGBB) (optional)
	TileOffset       Offset     `json:"tileoffset"`       // https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-tileoffset
	Grid             Grid       `json:"grid"`             // (Optional) https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-grid
	Properties       []Property `json:"properties"`       // A list of properties (name, value, type)
	Terrains         []Terrain  `json:"terrains"`         // Array of Terrains (optional)
	Tiles            []Tile     `json:"tiles"`            // Array of Tiles (optional)
	WangSets         []WangSet  `json:"wangsets"`         // Array of Wang sets (since 1.1.5)
}

// The Grid element is only used in case of isometric orientation,
// and determines how tile overlays for terrain and collision information are rendered.
type Grid struct {
	Orientation string `json:"orientation"` // "orthogonal" or "isometric"
	Width       int    `json:"width"`       // Width of a grid cell
	Height      int    `json:"height"`      // Height of a grid cell
}

// Tile TODO describe
type Tile struct {
	ID          int        `json:"id"`          // Local ID of the tile
	Type        string     `json:"type"`        // The type of the tile (optional)
	Properties  []Property `json:"properties"`  // A list of properties (name, value, type)
	Animation   []Frame    `json:"animation"`   // Array of Frames
	Terrain     []int      `json:"terrain"`     // Index of terrain for each corner of tile
	Image       string     `json:"image"`       // Image representing this tile (optional)
	ImageHeight int        `json:"imageheight"` // Height of the tile image in pixels
	ImageWidth  int        `json:"imagewidth"`  // Width of the tile image in pixels
	ObjectGroup Layer      `json:"objectgroup"` // Layer with type "objectgroup" (optional)
}

// Frame TODO describe
type Frame struct {
	Duration int `json:"duration"` // Frame duration in milliseconds
	TileID   int `json:"tileid"`   // Local tile ID representing this frame
}

// Terrain TODO describe
type Terrain struct {
	Name string `json:"name"` // Name of terrain
	Tile int    `json:"tile"` // Local ID of tile representing terrain
}

// WangSet TODO describe
type WangSet struct {
	CornerColors []WangColor `json:"cornercolors"` // Array of Wang colors
	EdgeColors   []WangColor `json:"edgecolors"`   // Array of Wang colors
	Name         string      `json:"name"`         // Name of the Wang set
	Tile         int         `json:"tile"`         // Local ID of tile representing the Wang set
	WangTiles    []WangTile  `json:"wangtiles"`    // Array of Wang tiles
}

// WangColor TODO describe
type WangColor struct {
	Color       string  `json:"color"`       // Hex-formatted color (#RRGGBB or #AARRGGBB)
	Name        string  `json:"name"`        // Name of the Wang color
	Probability float64 `json:"probability"` // Probability used when randomizing
	Tile        int     `json:"tile"`        // Local ID of tile representing the Wang color
}

// WangTile TODO describe
type WangTile struct {
	TileID int     `json:"tileid"` // Local ID of tile
	WangID [8]byte `json:"wangid"` // Array of Wang color indexes (uchar[8])
	DFlip  bool    `json:"dflip"`  // Tile is flipped diagonally
	HFlip  bool    `json:"hflip"`  // Tile is flipped horizontally
	VFlip  bool    `json:"vflip"`  // Tile is flipped vertically
}

// An ObjectTemplate is written to its own file
// and referenced by any instances of that template.
type ObjectTemplate struct {
	Type    string  `json:"type"`    // "template"
	Tileset Tileset `json:"tileset"` // External tileset used by the template (optional)
	Object  Object  `json:"object"`  // The object instantiated by this template
}
