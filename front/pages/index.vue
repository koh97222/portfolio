<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <v-row class="top">
        <div>
          <h1>Foodience</h1>
          <p>~Foodienceは、自炊する1人暮らしの方を応援する新しいSNSです~</p>
          <v-btn color="primary" outlined @click="showModal('新規登録')"
            >新規登録</v-btn
          >
          <v-btn color="primary" outlined @click="showModal('ログイン')"
            >ログイン</v-btn
          >
          <v-btn color="primary" outlined>お試しログイン</v-btn>
        </div>
      </v-row>
    </v-flex>

    <!-- ユーザー新規登録・ログインダイアログ -->
    <v-dialog v-model="dialog" persistent max-width="500px">
      <v-card>
        <v-card-title>
          <span class="headline">{{ mode }}</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="user.ID"
                  label="ID*"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="user.Password"
                  label="Password*"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
            <!-- 新規登録時 -->
            <v-row v-if="newRegist">
              <v-col cols="12">
                <v-text-field
                  v-model="user.ConfirmPassword"
                  label="Password* - 確認用 -"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="user.Name"
                  label="Name*"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="user.Email"
                  label="Email*"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="user.Year"
                  :items="year"
                  label="Year*"
                  required
                ></v-select>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="user.Month"
                  :items="month"
                  label="Month*"
                  required
                ></v-select>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="user.Day"
                  :items="['男性', '女性', 'その他', '54+']"
                  label="Day*"
                  required
                ></v-select>
              </v-col>
            </v-row>
            <!-- -->
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false"
            >Close</v-btn
          >
          <v-btn color="blue darken-1" text @click="dialog = false">{{
            mode
          }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- -->
  </v-layout>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      dialog: false,
      mode: null,
      user: {
        ID: null,
        Password: null,
        ConfirmPassword: null,
        Name: null,
        Email: null,
        Year: null,
        Month: null,
        Date: null,
        Sex: null,
      },
    }
  },
  computed: {
    year() {
      const YearList = []
      for (let i = 1920; i <= 2015; i++) {
        YearList.push(i)
      }
      return YearList
    },
    month() {
      const monthList = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
      return monthList
    },
    day() {
      const dayList = []
      return dayList
    },
    newRegist() {
      return this.mode === '新規登録'
    },
  },

  methods: {
    /**
     * ユーザーの新規登録を行う。
     * 登録OK→次画面へ遷移。
     * 登録NG→エラーメッセージを表示。
     */
    userNewRegist(user) {
      // 入力必須チェック
      const inputCheck = (user) => {
        user.forEach((info) => {
          if (!info) return false
        })
        return true
      }
      // パスワードチェック
      const validPassword = (user) => {
        if (user.Password === user.ConfirmPassword) {
          return true
        }
        return false
      }

      // -- 実処理スタート --
      if (!(inputCheck(user) || validPassword(user))) {
        // メッセージ
        return
      }

      // 登録処理
      this.$axios
        .post('http://localhost:8082/registUser')
        .then(() => {})
        .catch(() => {})
    },
    /**
     * モーダルの表示を行う。
     */
    showModal(mode) {
      this.dialog = !this.dialog
      this.mode = mode
    },
  },
}
</script>
