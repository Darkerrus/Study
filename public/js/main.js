const { createApp } = Vue;
createApp({
  data() {
    return {
      data: []
    };
  },
  computed: {
    
  },
  methods: {
    getDogs() {

    }
  },
  mounted() {
    this.data = document.querySelectorAll('[data-kind]')
    console.log(this.data);
  },
}).mount("#app");