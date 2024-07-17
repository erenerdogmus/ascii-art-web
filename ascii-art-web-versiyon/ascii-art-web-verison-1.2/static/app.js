document.getElementById('textarea').addEventListener('focus', function() {
    document.querySelector('.inputBx i').style.display = 'none';
});
function downloadFile() {
    var textToSave = document.querySelector('.inputBx pre').textContent;
    var hiddenElement = document.createElement('a');
    hiddenElement.href = 'data:attachment/' + ',' + encodeURI(textToSave);
    hiddenElement.target = '_blank';
    hiddenElement.download = 'ascii_art.';
    hiddenElement.click();
}
function clearValue() {
    document.getElementById('fileNameInput').value = '';
}