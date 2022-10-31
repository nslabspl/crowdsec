// Dont forget to adjust both message text and style to your likings/needs.
//
// General format of styling is:
// document.getElementById('cookieMessageDiv').style.<css-attribute> = <value>;
let cookieMessage = 'This webpage uses cookies.';
let addCMDiv = new function addToPage(){
	document.body.appendChild(document.createElement('div').id('cookieMessageDiv'));
	document.getElementById('cookieMessageDiv').innerHTML = cookieMessage;
	document.getElementById('cookieMessageDiv').style.height = auto;
	document.getElementById('cookieMessageDiv').style.backgroundColor = lightblue;
}