const express =require('express')
const app=express()
app.get('/',(req,res)=>{
    res.send("hello guys")
})
app.listen(3000,()=>{
    console.log("server has been started")
})