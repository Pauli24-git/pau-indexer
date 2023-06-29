const app = Vue.createApp({
    data() {
        return {
            templateContent: '',
            term:'',
            field:''
        };
    },
        methods: {
            fetchData() {
                urlbase='http://localhost:3080/search?'
                term= 'term=' +  this.term
                field= 'field=' + ((this.field == 'ALL') ? '_all' : this.field)
              fetch(urlbase + term + '&' + field )
                .then(response => response.text())
                .then(data => {
                  this.templateContent = data;
                })
                .catch(error => {
                  console.error(error);
                });
            }
          }
    }).mount('#app');

