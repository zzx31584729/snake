let root = document.getElementsByClassName("root");
root.item(0).addEventListener("click",function (el){
    if (el.target.style.backgroundColor === "black") {
        el.target.style.backgroundColor = "blue"
    }else{
        el.target.style.backgroundColor = "black"
    }
})