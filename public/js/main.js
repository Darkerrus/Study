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
          this.filteredData.push(el)
        }
      })
    },
    getCats() {
      toDefault()
      this.data.forEach(el => {
        if(el.dataset.kind === 'Cat') {
          this.filteredData.push(el)
        }
      })
    },
    getBird() {
      toDefault()
      this.data.forEach(el => {
        if(el.dataset.kind === 'Bird') {
          this.filteredData.push(el)
        }
      })
    }
  },
  mounted() {
    this.data = document.querySelectorAll('[data-kind]')
    this.pNode = document.getElementById('.animals');
  },
}).mount("#app");