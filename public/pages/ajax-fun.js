// Vanilla JavaScript AJAX
function loadDoc() {
    const xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            document.getElementById("vanilla-demo").innerHTML = this.response
        }
    };
    xhttp.open("GET", "https://cors-anywhere.herokuapp.com/http://carnes.cc/code/ajax_example.txt", true)
    xhttp.send();
}

// jQuery: AJAX = Asynchronous JavaScript And XML
function loadDoc2() {
    $("#jquery-demo").load("https://cors-anywhere.herokuapp.com/http://carnes.cc/code/ajax_example.txt")
}

function loadDoc3() {
    // .load() automatically loads data into an element
    $("#demo3").load("https://cors-anywhere.herokuapp.com/http://carnes.cc/code/ajax_example.txt",
        function (responseTxt, statusTxt, xhr) {
            const span = $('#demo3-addition')
            if (statusTxt === "success")
                span.text(" Attn: External content loaded successfully!")
            if (statusTxt === "error")
                span.text(" Error: " + xhr.status + ": " + xhr.statusText)
        });
}

function getDoc() {
    // .get() doesn't automatically load data into an element
    $.get("https://cors-anywhere.herokuapp.com/http://carnes.cc/code/ajax_example.txt",
        function (data, status) {
            console.log("Data: " + data + "\nStatus: " + status)
        })
}

$(function () {
    $('#jquery-btn').click(loadDoc2)
    $('#demo3-btn').click(loadDoc3)
    $('#demo-get-btn').click(getDoc)
});