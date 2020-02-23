

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
        page: 1,
        limit: 10,
        count: 0,
        list: []
    },

    created() {
        this.getList()
    },

    mounted() {
    },

    methods: {

        getList() {
            this.http.Get('/api/article/list', {
                page: this.page,
                    limit: this.limit
            })
                .then(function (res) {
                    if (res.code == 200) {
                        this.list = res.data.item
                        this.count = res.data.count
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });


        }
    },
})