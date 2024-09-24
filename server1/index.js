require('dotenv').config()
const express = require('express')
const server = express()
const bcrypt = require('bcrypt')
const users=[]
server.use(express.static(process.env.STATIC_FOLDER))
server.use(express.json())
server.use(express.urlencoded({extended:false}))
server.set('view-engine','ejs')
server.get("/home",(req,res)=>{
    res.render("index.ejs")
})
server.get("/register",(req,res)=>{
    res.render("register.ejs")
})
server.get("/login",(req,res)=>{
    res.render("login.ejs")
})
server.post("/register",async (req,res)=>{
    try{
        const hashedpassword = await bcrypt.hash(req.body.password,10)
        users.push({
            id:Date.now().toString(),
            name:req.body.name,
            email:req.body.name,
            password:hashedpassword
        })
        res.redirect("/login")


    }
    catch{
        res.redirect("/register")

    }
    console.log(users)
})
server.listen(process.env.PORT,()=>{
    console.log("server is running on",process.env.PORT)
})
