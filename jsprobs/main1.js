require("dotenv").config()
const express=require("express")
const app=express()
app.use(express.json())
const courses=[
    {id:1,name:"course1"},
    {id:2,name:"course2"}
]

app.get("/home",(req,res)=>{
    res.send("hello world")
})

//for getting all the courses
app.get("/courses",(req,res)=>{
    res.send(courses)
})
//for getting a particular course
app.get("/courses/:id",(req,res)=>{
    const course=courses.find(c=>c.id===parseInt(req.params.id))
    if(!course){
        res.status(400).send("cannot find the course")

    }
    res.send(course)
})
//for creating a course
app.post("/courses",  (req,res)=>{
    const course = {
        id:courses.length+1,
        name:req.body.name
    }
    courses.push(course)
    res.send(course)

})
//for updating a particular course
app.put("/courses/:id",(req,res)=>{
    const course=courses.find(c=>c.id===parseInt(req.params.id))
    if(!course){
        res.status(400).send("cannot find the course")

    }
    course.name=req.body.name
    res.send(course)

})

//deleting a particular course
app.delete("/courses/:id",(req,res)=>{
    const course=courses.find(c=>c.id===parseInt(req.params.id))
    if(!course){
        res.status(400).send("cannot find the course")
    }
    const index=courses.indexOf(course)
    courses.splice(index,1)
    res.send(course)
})
app.listen(process.env.PORT,()=>{
    console.log(`server has started on ${process.env.PORT}`)
})

