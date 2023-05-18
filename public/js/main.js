const { createApp } = Vue;
createApp({
  data() {
    return {
      data: [],
      filteredData: [],
      pNode: {}
    };
  },
  methods: {
    toDefault() {
      this.filteredData = []
      while (this.pNode.hasChildNodes()){
        this.pNode.removeChild(el.lastChild);
      }
    },
    getAll() {
      toDefault()
      this.filteredData = this.data
    },
    getDogs() {
      toDefault()
      this.data.forEach(el => {
        if(el.dataset.kind === 'Dog') {
          this.pNode.appendChild(el)
        }
      })
    },
    getCats() {
      toDefault()
      this.data.forEach(el => {
        if(el.dataset.kind === 'Cat') {
          this.pNode.appendChild(el)
        }
      })
    },
    getBird() {
      toDefault()
      this.data.forEach(el => {
        if(el.dataset.kind === 'Bird') {
          this.pNode.appendChild(el)
        }
      })
    }
  },
  mounted() {
    this.data = document.querySelectorAll('[data-kind]')
    this.pNode = document.getElementById('.animals');
  },
}).mount("#app");