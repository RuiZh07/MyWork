const ndef = new NFEFReader();

async function writeNFC(){
    try {
        await ndef.write({
            records: [{ recordType: "url", data: "https://themaverse.io/"}]
        });
    } catch {
        var text = document.getElementById('input1').value
        alert("Write failed, Please try again!"+ text )
    };
}

function getID(){
    var text=document.getElementById('input1').value;
    window.alert(text)
}

function test(){
    alert("something");
}