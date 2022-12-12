const ndef = new NFEFReader();

async function writeNFC(){
    try {
        await ndef.write({
            records: [{ recordType: "url", data: "https://themaverse.io/"}]
        });
    } catch {
        alert("Write failed, Please try again!")
    };
}

function getID(){
    var text=document.getElementById('input1').value;
    window.alert(text)
}

function test(){
    alert("something");
}