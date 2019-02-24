<template>
  <div>
    <md-field>
      <label>英単語を入力してください ※改行で区切る</label>
      <md-textarea v-model="words"></md-textarea>
    </md-field>
      <div class="md-layout-item md-layout md-gutter">
        <md-button @click="search" v-if="words.length" class="md-dense md-raised md-primary">変換開始！</md-button>
        <md-button @click="download" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">ダウンロード</md-button>
        <md-button @click="downloadMac" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">ダウンロード for Mac</md-button>
        <md-button @click="load" class="md-dense md-raised md-primary">再ロード</md-button>
        <md-button @click="clear" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">クリア</md-button>
      </div>
    <md-table class="wordTable">
      <md-table-row>
        <md-table-head>英単語</md-table-head>
        <md-table-head>発音記号</md-table-head>
        <md-table-head>意味</md-table-head>
      </md-table-row>
      <md-table-row v-for="word in response" v-bind:key="word.key" class="results">
        <md-table-cell>{{word.word}}</md-table-cell>
        <md-table-cell>{{word.pron}}</md-table-cell>
        <md-table-cell>{{word.meaning}}</md-table-cell>
      </md-table-row>
    </md-table>
    <!--
    <p>{{response}}</p>
    -->
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Vocabulary',
  data: () => ({
    words: '',
    response: [],
    dowwnloadDisplay: false,
    host: 'dumblepy.site'
  }),
  methods: {
    search: function () {
      // const url = 'https://' + this.host + '/products/vocabulary/api/searchByWords'
      // const url = 'https://dumblepy.site/products/vocabulary/api/searchByWords'
      const url = 'http://localhost:8000/products/vocabulary/api/searchByWords'
      const body = JSON.stringify({ words: this.words.split(/\n/) }) // 改行コードで分割して配列にする
      console.log(body)
      const headers = {
        'Accept': 'application/json'
      }

      axios.post(url, body, headers)
        .then(response => {
          this.response = response.data
          if (response.data.length > 0) {
            this.dowwnloadDisplay = true
            this.save(this.words, JSON.stringify(response.data))
          }
        })
        .catch(err => {
          console.error(err)
        })

      console.log(this.response)
    },
    download: function () {
      const url = 'http://' + this.host + '/products/vocabulary/api/downloadCSV'
      const method = 'POST'
      const body = JSON.stringify({ request: this.response })
      console.log(body)
      const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }

      fetch(url, { method, headers, body })
        .then(response => {
          // if (response.text.length > 0) {
          return response.blob()
          // }
        })
        .then(response => {
          let anchor = document.createElement('a')
          anchor.download = 'vocabulary.csv'
          anchor.href = window.URL.createObjectURL(response)
          anchor.click()
        })
    },
    downloadMac: function () {
      const url = 'http://' + this.host + '/rubyapi/downloadCSVMac'
      const method = 'POST'
      const body = JSON.stringify({ request: this.response })
      const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }

      fetch(url, { method, headers, body })
        .then(response => {
          // if (response.text.length > 0) {
          return response.blob()
          // }
        })
        .then(response => {
          let anchor = document.createElement('a')
          anchor.download = 'vocabulary.csv'
          anchor.href = window.URL.createObjectURL(response)
          anchor.click()
        })
    },
    clear: function () {
      this.response = []
      this.dowwnloadDisplay = false
    },
    save: function (words, response) {
      window.localStorage.setItem('words', words)
      window.localStorage.setItem('response', response)
    },
    load: function () {
      this.words = window.localStorage.getItem('words')
      this.response = JSON.parse(window.localStorage.getItem('response'))
      this.dowwnloadDisplay = true
    }
  },
  created: function () {
    if (window.localStorage.getItem('response')) {
      this.load()
    }
    let host = location.host
    // if (host === 'localhost:8080') {
    //   host = 'localhost:3000'
    // }
    this.host = host
    console.log(host)
  }

}
</script>

<style scoped>
  .wordTable {
    margin-top: auto !important;
  }
</style>
