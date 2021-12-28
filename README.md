Hi Github. I just learn how to use tree on Golang.

Here is a program to print all the component from an item. Here the spec :
    - Each item can has either component(s) or part(s)
    - Each part can has either component(s) or part(s)
    - Component is the smallest things

U can add this following LOC on main function :
Define your item
    - itemVar := defineItem(itemName)

Insert part to an item
    - itemVar := itemVar.insertPart(partName)

Insert part to an part
    - childPart := parentPart.insertPart(partName)

Insert component to an part
    - partVar.insertComponent(partName)

Kindly contribute if you have idea to make this program more user friendly. Diagram Example that implemented on this code: 