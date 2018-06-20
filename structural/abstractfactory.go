package structural

// Abstract Factory pattern is used to create families of related or dependent objects without specifying their concrete classes.
// It is also known as Kit.

// Where can you apply this pattern -
// 	 Where the system is independent of how its products are created, composed and represented.
//   Also it should be configured with one of multiple family of products.
//   The family of related products are designed to be used together.
//   You want to explose interface of the products not their concrete implementation.

// Benefits of using this pattern -
//   1. It isolates concrete classes as the factory encapsulates the creation of objects. They are controlled by common interfaces.
//
//   2. It makes exchange of families of products easier. Only the instantiation of products changes in factory implementation.
//
//   3. It ensures consistency among products as only one family of products are initialzed at one time.

// The only problem using this pattern is that extending the factory to create more objects is difficult.
