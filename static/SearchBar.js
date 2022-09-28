const searchBar = document.body.querySelector("#searchBar");

searchBar.addEventListener("keyup", (e) => {
    const searchedLetters = e.target.value;
    const cards = document.querySelectorAll(".card");
    filterElements(searchedLetters, cards)
});

function filterElements(letters, elements) {
    for (let i=0; i<elements.length; i++){
        if (elements[i].textContent.toLowerCase().includes(letters)) {
            elements[i].style.display = "block"
        } else {
            elements[i].style.display = "none";
        }
    }
}

var OneChoise = true;
//        <input type="submit" id="SearchBtn" name="Filter" value="SearchBtn">

function MultipleChoiseOn() {
    const MultipleChoiseBox = document.getElementById("MultipleChoise");
    const SearchBtn = document.getElementById("SearchBtn");
    if (MultipleChoiseBox.checked) {
        SearchBtn.style.display = "block";
        OneChoise = false;
    } else {
        SearchBtn.style.display = "none";
        OneChoise = true;
    }
}

function filterSub(ClickedBox) {
    console.log(ClickedBox);
    var BoxVide = "BoxVide";
    var Box = document.getElementById(ClickedBox);
    
    if (Box.checked) {
        window.sessionStorage.setItem("ClickedBox", ClickedBox);
        var att = document.createAttribute("checked");
        Box.setAttributeNode(att);

    } else {
        window.sessionStorage.setItem("ClickedBox", BoxVide);
    }

    if (OneChoise) {
        var AllBox = document.getElementsByName("Filter");
        for (let i=0; i<AllBox.length; i++){
            if (AllBox[i].checked && AllBox[i] != Box) {
                AllBox[i].checked = false;
            }
        }
    
        
        document.getElementById("FilterForm").submit();
    }
    

}
window.onload = LoadFunc;

function LoadFunc() {
    var ClickedBox = window.sessionStorage.getItem("ClickedBox");
    if (ClickedBox != null && ClickedBox != "BoxVide" ) {
        console.log(ClickedBox);
        var Box = document.getElementById(ClickedBox);
        if (Box != null) {
            Box.checked = true;
        }
        
    }
    
}
