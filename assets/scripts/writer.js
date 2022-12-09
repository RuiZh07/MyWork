const ndef = new NFEFReader();
try {
    await ndef.write({
        records: [{ recordType: "url", data: "https://themaverse.io/"}]
    });
} catch {
    console.log("Write failed, Please try again!")
};