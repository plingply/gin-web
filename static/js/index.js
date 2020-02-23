

drawbackground({
    zindex: -1,//canvas 定位 
    opacity: 0.5,//canvas 透明度 
    color: "0.0.0",//画线颜色 
    count: 30//点的数量 
})

Vue.prototype.http = Axios

var app = new Vue({
    el: "#app",

    data: {

        tabIndex: "",

        page: 1,
        limit: 10,
        count: 0,
        list: [],
        type: []
    },

    created() {
        this.getTypeList()
        this.getList()
    },

    mounted() {
        
    },

    methods: {

        tabClick(code){
            this.tabIndex = code 
            this.getList()
        },

        getList() {
            this.http.Get('/api/article/list', {
                page: this.page,
                limit: this.limit,
                type: this.tabIndex
            })
                .then( (res)=> {
                    if (res.code == 200) {
                        this.list = res.data.item
                        this.count = res.data.count
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });


        },

        getTypeList(){
            this.http.Get("/api/article/type")
            .then( (res)=> {
                if (res.code == 200) {
                    this.type = res.data
                }
            })
            .catch(function (error) {
                console.log(error);
            });
        },
    }
})