const data = document.querySelectorAll('[data-kind]')
const pNode = document.getElementById('.animals');

const toDefault = () => {
  while (pNode.hasChildNodes()){
    pNode.removeChild(pNode.lastChild);
  }
}

// const getAll = () => {
//   toDefault()
// }

const getDogs = document.querySelector('#getDogs')
const getCats = document.querySelector('#getCats')
const getBirds = document.querySelector('#getBirds')

getDogs.onclick = () => { 
  toDefault()
  data.forEach(el => {
    if(el.dataset.kind === 'Dog') {
      pNode.appendChild(el)
    }
  })
}

getDogs.onclick = () => { 
  toDefault()
  data.forEach(el => {
    if(el.dataset.kind === 'Dog') {
      pNode.appendChild(el)
    }
  })
}

getCats.onclick = () => { 
  toDefault()
  data.forEach(el => {
    if(el.dataset.kind === 'Cat') {
      pNode.appendChild(el)
    }
  })
}

getBirds.onclick = () => { 
  toDefault()
  data.forEach(el => {
    if(el.dataset.kind === 'Bird') {
      pNode.appendChild(el)
    }
  })
}