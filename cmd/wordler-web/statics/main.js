function getInputStorageKey(key) {
    return "INPUT_" + key;
}

Vue.component('component-filter', {
    props: ['field', 'name', 'value', 'hint'],
    template: "#component-filter",
    methods: {
        store: function (value) {
            localStorage.setItem(getInputStorageKey(this.$props.field), value);
        }
    },
})

var app = new Vue({
    el: '#app',
    data: {
        filters: [
            { 'key': 'has', 'name': 'Has', 'data': '', 'hint': 'a,b,c' },
            { 'key': 'no', 'name': 'No', 'data': '', 'hint': 'x,y,z' },
            { 'key': 'at', 'name': 'At', 'data': '', 'hint': '0:a,1:b' },
            { 'key': 'na', 'name': 'Not at', 'data': '', 'hint': '2:c' },
        ],
        result: '',
    },
    mounted() {
        for (let i = 0; i < this.filters.length; i++) {
            let value = localStorage.getItem(getInputStorageKey(this.filters[i].key));
            if (value) {
                this.filters[i].data = String(value);
            }
        }
    },
    methods: {
        run: function () {
            payload = {}
            for (filter of this.filters) {
                payload[filter.key] = filter.data.split(',').filter(s => s.length > 0);
            }
            this.result = '';
            this.sendFilter(this.addResult, this.addResult, payload);
        },
        clearInput: function () {
            for (let i = 0; i < this.filters.length; i++) {
                this.filters[i].data = '';
                localStorage.removeItem(getInputStorageKey(this.filters[i].key));
            }
        },
        addResult: function (data) {
            this.result += data;
            this.result += "\n";
        },
        sendFilter: function (succ, fail, data) {
            fetch(
                `./filter`, {
                body: JSON.stringify(data),
                method: 'POST',
                headers: {
                    'content-type': 'application/json',
                },
            }).then(function (res) {
                if (res.status !== 200) {
                    fail('Failed with Status Code: ' + res.status);
                    res.text().then(fail);
                } else {
                    res.text().then(succ);
                }
            }).catch(function (err) {
                console.log('Fetch Error :-S', err);
                fail(err);
            });
        }
    }
})
