function resolve(url) {
	console.log('check');
	window.location.href = url;
}

function display(value) {
	displayNone();
	if (value == 'personal') {
		document.getElementById("section-personal").style.display = "block";
		document.getElementById('link-personal').classList.add("active");
	}
	else if (value == 'contact') {
		document.getElementById("section-contact").style.display = "block";
		document.getElementById('link-contact').classList.add("active");
	}
	else {
		document.getElementById("section-history").style.display = "block";
		document.getElementById('link-history').classList.add("active");
	}
}

function displayNone() {
	document.getElementById("section-personal").style.display = "none";
	document.getElementById("section-contact").style.display = "none";
	document.getElementById("section-history").style.display = "none";
	document.getElementById('link-personal').classList.remove("active");
	document.getElementById('link-contact').classList.remove("active");
	document.getElementById('link-history').classList.remove("active");
}

function upload() {
	document.getElementById('picture').click();
}

function remove() {
	confirm("Do you want to remove profile picture?");
}