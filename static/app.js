var convApp = new Vue({
    el: '#convApp',
    data: {
        curency: [
            "USD",
            "EUR",
            "GBP",
            "AUD",
            "AZN",
            "AMD",
            "BGN",
            "HUF",
            "RUR",
        ],
        selected_user_curency: "RUR",
        selected_out_curency: "USD",
        amount: "1",
        result_conv: null,
    },
    methods: {
        // проверяем пришла ли валюта пользователя
        select_user_curency: function () {
            console.log('валюта', this.selected_user_curency);
        },
        // проверяем пришла ли выходная валюта
        select_out_curency: function () {
        console.log('валюта на выходе', this.selected_out_curency);
        },
        convertator_cur: function () {
            if (this.amount != "" && this.selected_user_curency!=this.selected_out_curency) {
                console.log('все ок');
                this.$http.post('http://localhost:9090/getRBK', { selected_user_curency: this.selected_user_curency, amount: this.amount, selected_out_curency: this.selected_out_curency }).then(response => {
                    this.result_conv = response.body.result;
                }, response => {
                    console.log("error:", response)
                });
            }
            else {
                alert("Заполните все поля")
            }
        }
    }
})


var login = new Vue({
    el: '#loginApp',
    data: {
        usernameLogin: null,
        username: null,
        id: null,
    },
    methods: {

        login: function () {
            if (this.usernameLogin) {
                this.$http.post('http://localhost:9090/auth', { username: this.usernameLogin }).then(response => {
                    this.username=response.body.user.username
                    this.id=response.body.user.id
                    console.log("user:", this.username,this.id)
                    window.location.href = '/pageConv' ;
                }, response => {
                    console.log("error:", response)
                });
            }
            else {
                alert("Введите имя пользователя!")
            }
        },
    },
})