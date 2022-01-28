'use strict';

function loadWasm(path) {
 const go = new Go();

 return new Promise((resolve, reject) => {
   WebAssembly.instantiateStreaming(fetch(path), go.importObject)
   .then(result => {
     go.run(result.instance);
     resolve(result.instance);
   })
   .catch(error => {
     reject(error);
   })
 })
}

function onLoad() {
  // This is a polyfill for FireFox and Safari
  if (!WebAssembly.instantiateStreaming) { 
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    }
  }

  loadWasm('main.wasm').then(wasm => {
    console.log('main.wasm is loaded 👋');
    handlePasswordUpdate();
  }).catch(error => {
    console.log("ouch", error);
  });

  const inputs = document.getElementsByTagName('input');
  for(let i=0; i < inputs.length; i++) {
    inputs[i].addEventListener('change', handlePasswordUpdate);
  }

  const rangeInput = document.getElementById('length-range');
  rangeInput.addEventListener('change', handleRangeLabelUpdate)
}

function handleRangeLabelUpdate(event) {
  const length = event.currentTarget.value;
  document.getElementById('length-label').textContent = `Length: ${length}`
}

function handlePasswordUpdate() {
  const length =  parseInt(document.getElementById('length-range').value);
  const includeMixedCase = document.getElementById('mixed-case-check').checked;
  const includeNumbers = document.getElementById('numbers-check').checked;
  const includeSpecial = document.getElementById('special-check').checked;
  
  document.getElementById('password-output').textContent =
    generatePassword(length, { includeMixedCase, includeNumbers, includeSpecial });
}

document.addEventListener("DOMContentLoaded", onLoad);
