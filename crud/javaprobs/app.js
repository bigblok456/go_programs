
function Person(id,name,dob){
    
    this.id=id
    this.name=name
    this.dob=new Date(dob)


}
Person.prototype.Getname=function(){
    return `${this.name}`
}
const person1 =  new person('1','kausik','3-4-1980')
console.log(person1.Getname())