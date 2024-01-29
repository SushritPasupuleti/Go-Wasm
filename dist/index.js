const go = new Go();

WebAssembly.instantiateStreaming(fetch('main.wasm'),
	go.importObject).then((result) => {
		go.run(result.instance);

		console.log("Result:", add(2, 3)); // call the 'add' function defined in the Go program
	});

window.addEventListener('DOMContentLoaded', function() {
	console.log("DOM loaded");
	document.getElementById("add-btn").addEventListener("click", onAdd);
});

function onAdd() {
	const a = document.getElementById("a").value;
	const b = document.getElementById("b").value;

	const result = add(parseInt(a), parseInt(b));
	console.log("Result:", result);
	document.getElementById("result").innerHTML = "<b>Result: </b>" + parseInt(result);
}
