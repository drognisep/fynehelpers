# Fyne Helpers

> **Note:**
> As much as I would enjoy continuing to work on this, it's based on Fyne.
> Fyne (among others) has made some great strides toward bringing GUI possibilities to Go, but in my opinion it has a long way to go to be production ready.
> 
> I'll be excited when that happens, but in the meantime I will be marking this repo as archived.
> I hope this will help to clarify the level of time I'm willing to devote to Fyne-based projects at this time.

I've been playing with Fyne for a bit, and I've found myself rewriting some of the same code a few different times.
Instead of continuing with this duplication, I'd rather bring these common tools together into one place.

This also includes a CLI for generating Fyne structures that are a bit tedious to write over and over.
There aren't a lot right now, but I'll be adding to the code gen capabilities as I go.

## Code Generation
To install the `fynehelper` command which provides code generation capabilities:
```shell
go install github.com/drognisep/fynehelpers/cmd/fynehelper@latest
```

See the current set of options by running `fynehelper --help` or `fynehelper generate --help` for generation options.

### Tree generation
Currently, the only generator available is to generate a Fyne v2 tree structure.
* Allows registering tap (single, double, secondary) handlers on tree nodes.
* Maintaining a central store of tree data.
* Dynamically adding/removing tree nodes, even from tap handlers.
* Walk the current state of the tree (read-only).
* Plug your data into the tree structure using your icon and/or text of choice using a consistent interface that works for any generated tree.

#### Example
```shell
fynehelper generate tree --pkg=testing --file=myTree --type=TestTree --event-tapped
```
This will generate `myTreeNode.go` and `myTreeTree.go` in the current directory for the custom node and tree code, respectively.
Their package will be set to testing, and the node will respond to single tap events.
The generated tree will be called `TestTreeTree` and has an accompanying constructor function, `NewTestTreeTree`.
Nodes are not meant to be interacted with directly (in this case it'll be `testTreeNode`), but will be managed by the tree.

#### Responding to events
Event handler functions may be set on the generated tree itself.
Enabling single tap handling will provide a `OnTapped` field which will receive the event from Fyne *and* contextual details about the tapped node's data.

## Packages

### layouthelp

This package is used for common use cases for custom layouts or widget layouts.

### stylehelp

This package contains ways to tweak the visual styling of canvas objects.

### testhelp

Provides testing helpers.

### generation

This package provides some base code used by generated structures, like the tree generator.
Because of this, generated code has a dependency on this package, so it should be added to projects using generated structures.
