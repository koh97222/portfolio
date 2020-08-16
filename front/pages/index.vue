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
                  v-model="input.UserID"
                  label="ID*"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  v-model="input.Password"
                  label="Password*"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
            </v-row>

            <!-- 新規登録時に表示 -->
            <v-row v-if="isNewRegist">
              <v-col cols="12">
                <v-text-field
                  v-model="input.ConfirmPassword"
                  label="Password* - 確認用(for confirmation) -"
                  type="password"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="6">
                <v-text-field
                  v-model="input.Name"
                  label="名前 Name*"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="6">
                <v-select
                  v-model="input.Sex"
                  :items="Object.values(sex)"
                  label="性別 Sex*"
                  required
                ></v-select>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  v-model="input.Email"
                  label="メールアドレス Email*"
                  required
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="input.Year"
                  :items="year"
                  label="生年月日 Year*"
                  required
                ></v-select>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="input.Month"
                  :items="month"
                  label="月 Month*"
                  required
                ></v-select>
              </v-col>
              <v-col cols="12" sm="4">
                <v-select
                  v-model="input.Day"
                  :disabled="isDayColumnDisabled"
                  :items="day"
                  label="日 Day*"
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
            >閉じる</v-btn
          >
          <v-btn color="blue darken-1" text @click="userNewRegist(input)">{{
            mode
          }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- -->
  </v-layout>
</template>

<script>
const userInitData = {
  UserID: null,
  Password: null,
  ConfirmPassword: null,
  Name: null,
  Email: null,
  Year: null,
  Month: null,
  Day: '',
  Sex: null,
}
export default {
  components: {},
  data() {
    return {
      dialog: false,
      mode: null,
      input: userInitData,
      errMsg: [],
    }
  },
  computed: {
    year() {
      const YearList = []
      // とりあえず1920年-2015年
      for (let i = 1920; i <= 2015; i++) {
        YearList.push(i)
      }
      return YearList
    },
    month() {
      const monthList = Array(12)
        .fill()
        .map((_, i) => i + 1)
      return monthList
    },
    day() {
      /**
       * isLeapYear うるう年がどうかを判定。
       * 【うるう年の条件】
       * ・西暦年が4で割り切れる年はうるう年
       * ・ただし、西暦年が100で割り切れる年は平年（うるう年ではない）
       * ・ただし、西暦年が400で割り切れる年はうるう年
       * ・※vue2-datepickerライブラリを使用すべきか悩んだが、あえて使わずに実装してみる。
       */
      const isLeapYear = (year) => {
        if (year % 400 === 0) {
          return true
        }
        if (year % 100 === 0) {
          return false
        }
        return year % 4 === 0
      }

      // is30Days 月の日数が30日かどうかを判定。
      const is30Days = (month) => {
        const list = [4, 6, 9, 11]
        return list.includes(month)
      }

      /**
       * dayMaker 日付のプルダウンを生成する関数。
       */
      const dayPullDownMaker = (input) => {
        let days = null
        if (input.Month === 2) {
          days = isLeapYear(input.Year) ? 29 : 28
        } else {
          days = is30Days(input.Month) ? 30 : 31
        }
        const dayList = Array(days)
          .fill()
          .map((_, i) => i + 1)
        return dayList
      }

      // 月が入力されている場合、日付のプルダウンを生成して返却します。
      return this.input.Month ? dayPullDownMaker(this.input) : null
    },

    // isDayColumnDisabled 年と月が未入力の場合、日付入力欄を非活性にします。
    isDayColumnDisabled() {
      return !(this.input.Year && this.input.Month)
    },
    // isNewRegist 新規登録であるかどうかを判定。
    isNewRegist() {
      return this.mode === '新規登録'
    },
    sex() {
      return ['男性', '女性', 'その他']
    },
  },

  watch: {
    // eslint-disable-next-line object-shorthand
    'input.Year': function () {
      this.input.Day = null
    },
    // eslint-disable-next-line object-shorthand
    'input.Month': function () {
      this.input.Day = null
    },
    // eslint-disable-next-line object-shorthand
    dialog: function () {
      // 入力情報をリセットする。
      this.input = JSON.parse(JSON.stringify(userInitData))
    },
  },

  methods: {
    /**
     * userNewRegist ユーザーの新規登録を行う。
     * 登録OK→次画面へ遷移。
     * 登録NG→エラーメッセージを表示。
     */
    userNewRegist(input) {
      // isAllInput 入力必須チェック
      const isAllInput = (input) => {
        let isValid = true
        Object.values(input).forEach((val) => {
          if (!val) {
            isValid = false
          }
        })
        return isValid
      }
      // validPassword パスワードチェック
      const validPassword = (input) => {
        // TODO: 3種類8文字以上のパスワードが入力されているかチェックする。
        if (input.Password && input.Password === input.ConfirmPassword) {
          return true
        }
        return false
      }

      // validEmail メールアドレスチェック
      const validEmail = (input) => {
        // TODO: 正規表現で、メールアドレスの形式かどうかチェックする。
        console.log(input)
        return true
      }

      // inputValid 入力された情報がすべて妥当か判断する関数。
      const inputValid = () => {
        return isAllInput(input) && validPassword(input) && validEmail(input)
      }
      // -- 実処理スタート --
      if (!inputValid()) {
        // eslint-disable-next-line no-console
        console.log('input err')
        return
      }

      // 登録処理
      this.$axios
        .post('http://localhost:8082/registUser', input)
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
