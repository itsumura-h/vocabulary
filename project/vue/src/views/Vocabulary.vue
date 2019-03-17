<template>
  <div>
    <md-field>
      <label>英単語を入力してください ※改行で区切る</label>
      <md-textarea v-model="words"></md-textarea>
    </md-field>
      <div class="md-layout-item md-layout md-gutter">
        <md-button @click="search" v-if="words.length" class="md-dense md-raised md-primary">変換開始！</md-button>
        <md-button @click="download" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">ダウンロード</md-button>
        <!--<md-button @click="downloadMac" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">ダウンロード for Mac</md-button>-->
        <md-button @click="load" class="md-dense md-raised md-primary">再ロード</md-button>
        <md-button @click="clear" v-show="dowwnloadDisplay" class="md-dense md-raised md-primary">クリア</md-button>
      </div>
    <md-table class="wordTable">
      <md-table-row>
        <md-table-head>英単語</md-table-head>
        <md-table-head>発音記号</md-table-head>
        <md-table-head>意味</md-table-head>
      </md-table-row>
      <md-table-row v-for="word in response" v-bind:key="word.Key" class="results">
        <md-table-cell>{{word.Word}}</md-table-cell>
        <md-table-cell>{{word.Pron}}</md-table-cell>
        <md-table-cell>{{word.Meaning}}</md-table-cell>
      </md-table-row>
    </md-table>
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
      const url = 'https://dumblepy.site/products/vocabulary/api/searchByWords'
      // const url = 'http://localhost:8000/products/vocabulary/api/searchByWords'
      const body = JSON.stringify({ words: this.words.split(/\n/) }) // 改行コードで分割して配列にする
      const headers = {
        'Accept': 'application/json',
        'mode': 'no-cors'
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
    },
    download: function () {
      const response = JSON.parse(window.localStorage.getItem('response'))
      let csvBody = [['id', '英単語', '発音', '意味']]
      for (let row of response) {
        let rowBody = [row.Key + 1, row.Word, row.Pron, row.Meaning]
        csvBody[row.Key + 1] = rowBody
      }

      let data = csvBody.map((row) => row.join(',')).join('\r\n')
      let bom = new Uint8Array([0xEF, 0xBB, 0xBF])
      let blob = new Blob([bom, data], { type: 'text/csv' })

      let anchor = document.createElement('a')
      anchor.download = 'vocabulary.csv'
      anchor.href = window.URL.createObjectURL(blob)
      anchor.click()
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
  }

}
</script>

<style scoped>
  .wordTable {
    margin-top: auto !important;
  }
</style>
