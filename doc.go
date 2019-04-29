// Using the defined wrappers, one can at least mask and use untyped APIs with gobind/gomobile,
// to avoid API exclusion or compile failures by paying the price of verbosity and performance.
//
// There are no getters for incompatible values, because gomobile will discard them anyway.
// On the other hand, a method would allow the introduction of a (duck typing) interface which will introduce more
// problems, because the gomobile generated proxy type would not have that method anymore and therefore cannot
// satisfy the required interface, leading to weired compilation errors.
package std
