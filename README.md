Hi Github. I just learn how to implement tree data type on Golang.

Here is a program to print all the component from an item. Here the spec :

    - Each item can has either component(s) or part(s)
    
    - Each part can has either component(s) or part(s)
    
    - Component is the smallest thing
    

U can add this following LOC on main function :

Define your item

`itemVar := defineItem(itemName)`

Insert part to an item

`itemVar := itemVar.insertPart(partName)`

Insert part to an part

`childPart := parentPart.insertPart(partName)`

Insert component to an part

`partVar.insertComponent(partName)`
