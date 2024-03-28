package microsoft

// ```ts
// model AModel {
//     b: BModel,
//     c: CModel,
// }
//
// model BModel {
//     c: CModel,
// }
//
// model CModel {
//     name: string,
// }
//
// op test(a: AModel): void;

// op test(a: AModel): CModel

// model CModel {
//   c: AModel
// }
// ```
//
// // Print all the usage paths of model `CModel` start from an operation "test". two paths:
// // "test" op -> "a" request -> AModel Model -> b ModelProperty -> BModel Model -> c ModelProperty -> CModel Model
// // "test" op -> "a" request -> AModel -> c -> CModel
// // "test" op -> "CModel" response -> "CModel" Model

// // Handle the cycle ref
// ```ts
// [
// {name: "test"
// type: Operation},
//
// {name: "a"
// type: Request},
//
// {name: "AModel"
// type: Model},
//
// {name: "b"
// type: ModelProperty},
//
// {name: "BModel"
// type: Model},
//
// {name: "c"
// type: ModelProperty},
//
// {name: "CModel"
// type: Model}
// ]
// ```
//
//
//
// ```java
// class Model {
//     private List<ModelProperty> properties;
//     public Model(List<ModelProperty> properties);
//     public List<ModelProperty> getProperties();
// }
//
// class Operation {
//     private String name;
//     private Request request;
//     private Response response;
//     public Operation(String name, Request request, Response response);
//     public String getName();
//     public Request getRequest();
//     public Response getResponse();
// }
//
// class ModelProperty {
//     private String name;
//     private Model modelType;
//     public ModelProperty(String name, Model modelType);
//     public String getName();
//     public Model getModelType();
// }
//
// class Request {
//     private String parameterName;
//     private Model type;
//     // ctor, getters
// }
//
// class Response {
//     private Model type;
//     // ctor, getters
// }
//
//
// class Node {
//     public String name;
//     public String type; // type can be: Operation | Request | Response | Model | ModelProperty
// }
//
// List<List<Node>> getUsagePaths(Operation operation, Model model);
// ```

type Type int

const (
	OperationType Type = iota
	RequestType
	ResponseType
	ModelType
	ModelPropertyType
)

func (t Type) String() string {
	// TODO: implement it
	return ""
}

// Protect ENUM

type Model interface {
	GetName() string
	GetProperties() []ModelProperty
}

type ModelProperty interface {
	GetName() string
	GetModel() Model
}

type Operation interface {
	GetName() string
	GetRequest() Request
	GetResponse() Response
}

type Request interface {
	GetParameterName() string
	GetModel() Model
}

type Response interface {
	GetModel() Model
}

// To support the cycle ref handling
// --> Gen hashmap key
func CmpModel(ref, target Model) bool {
	// TODO: implement
	return false
}

type Node struct {
	Name string
	// type can be: Operation | Request | Response | Model | ModelProperty
	Type Type
}

// Model -> multiple children
// Operation -> 2 children
// ModelProperty, Request, Response -> 1 child

// op test(a: AModel): CModel
//
// // Print all the usage paths of model `CModel` start from an operation "test". two paths:
// // "test" op -> "a" request -> AModel Model -> b ModelProperty -> BModel Model -> c ModelProperty -> CModel Model
// // "test" op -> "a" request -> AModel -> c -> CModel
// // "test" op -> "CModel" response -> "CModel" Model
func GetUsagePaths(o Operation, target Model) [][]Node {
	// Consider using DFS over BFS
	currentPaths := make([]Node, 0, 10)
	usagePaths := make([][]Node, 0, 10)
	dfs(o, target, currentPaths, &usagePaths)
	return usagePaths
}

func dfs(input any, target Model, currentPaths []Node, usagePaths *[][]Node) {
	if o, ok := input.(Operation); ok {
		node := Node{
			Name: o.GetName(),
			Type: OperationType,
		}
		currentPaths = append(currentPaths, node)
		cloneCurrentPaths := make([]Node, len(currentPaths))
		copy(cloneCurrentPaths, currentPaths)
		if o.GetRequest() != nil {
			dfs(o.GetRequest(), target, currentPaths, usagePaths)
		}
		if o.GetRequest() != nil {
			dfs(o.GetResponse(), target, cloneCurrentPaths, usagePaths)
		}
	} else if r, ok := input.(Request); ok {
		node := Node{
			Name: r.GetParameterName(),
			Type: RequestType,
		}
		currentPaths = append(currentPaths, node)
		dfs(r.GetModel(), target, currentPaths, usagePaths)
	} else if r, ok := input.(Response); ok {
		node := Node{
			Name: r.GetModel().GetName(),
			Type: RequestType,
		}
		currentPaths = append(currentPaths, node)
		dfs(r.GetModel(), target, currentPaths, usagePaths)
	} else if m, ok := input.(Model); ok {
		node := Node{
			Name: m.GetName(),
			Type: ModelType,
		}
		currentPaths = append(currentPaths, node)
		if CmpModel(m, target) {
			*usagePaths = append(*usagePaths, currentPaths)
			return
		}
		for _, p := range m.GetProperties() {
			cloneCurrentPaths := make([]Node, len(currentPaths))
			copy(cloneCurrentPaths, currentPaths)
			dfs(p, target, cloneCurrentPaths, usagePaths)
		}
	} else if p, ok := input.(ModelProperty); ok {
		node := Node{
			Name: p.GetName(),
			Type: ModelPropertyType,
		}
		currentPaths = append(currentPaths, node)
		dfs(p.GetModel(), target, currentPaths, usagePaths)
	}
}

// CModel
// A -> C -> A
// A -> B -> A -> C
