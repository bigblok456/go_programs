const num=[{
    id:1,
    name:"kausik",
    lastname:"thummala",
    iscomp:true
},
{
    id:2,
    name:"vulli",
    lastname:"abhinav",
    iscomp:true
},{
    id:3,
    name:"kausik",
    lastname:"bigblok",
    iscomp:false
},]
const numJSON = JSON.stringify(num)
//converting it into JSON so as to transfer data among web applications
console.log(numJSON)
const num1={
    id:2,
    firstname:"kausik",
    lastname:"thummala",
    address:{
        state:"telangana",
        city:"hyderabad"

    },
    books:["nice", "killa", "main"]
}
//num1 is an object literal,it has key-value pairs
//
//looping and higher order methods
//foreach,map
num.forEach(function(n){
    console.log(n.name)
})
const num2=num.map(function(num){
    return num.name
})
console.log(num2)
const num3=num.filter(function (num){
    return num.iscomp===true
}).map(function (num1){
    return num1.lastname
})
console.log(num3)


