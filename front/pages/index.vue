<template>
  <!-- <v-layout column justify-center align-center> -->
  <v-layout>
    <v-row>
      <v-row class="top" style="margin: 10px;">
        <v-col cols="12" sm="6" md="6" xs="6"> </v-col>
        <v-col cols="12" sm="6" md="6" xs="6">
          <v-layout justify-center align-center mt-5 mb-5>
            <div>
              <h1>Foodience</h1>
              <p>~Foodienceは、自炊する1人暮らしの方を応援する新しいSNSです~</p>
              <v-btn color="primary" outlined @click="showModal('新規登録')"
                >新規登録</v-btn
              >
              <v-btn color="primary" outlined @click="showModal('ログイン')"
                >ログイン</v-btn
              >
              <v-btn color="primary" outlined @click="trialLogin()"
                >お試しログイン</v-btn
              >
            </div>
          </v-layout>
        </v-col>
        <v-row class="contents mt-150">
          <v-layout justify-center align-center
            ><div>
              <h1 style="text-align: center;">機能紹介</h1>
              <v-row>
                <div style="margin: 5px;">
                  <v-card
                    class="mx-auto my-12"
                    elevation="2"
                    outlined
                    shaped
                    tile
                    max-width="374"
                    style="height: 600px; width: 400px;"
                  >
                    <v-img
                      height="250"
                      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
                    ></v-img>
                    <v-card-title>
                      レシピ投稿
                    </v-card-title>
                  </v-card>
                </div>
                <div style="margin: 5px;">
                  <v-card
                    class="mx-auto my-12"
                    elevation="2"
                    outlined
                    shaped
                    tile
                    max-width="374"
                    style="height: 600px; width: 400px;"
                  >
                    <v-img
                      height="250"
                      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
                    ></v-img>
                    <v-card-title>
                      レシピ検索
                    </v-card-title>
                  </v-card>
                </div>
                <div style="margin: 5px;">
                  <v-card
                    class="mx-auto my-12"
                    elevation="2"
                    outlined
                    shaped
                    tile
                    max-width="374"
                    style="height: 600px; width: 400px;"
                  >
                    <v-img
                      height="250"
                      src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
                    ></v-img>
                    <v-card-title>
                      Foodiest ranking
                    </v-card-title>
                  </v-card>
                </div>
              </v-row>
            </div>
          </v-layout>
        </v-row>
      </v-row>
    </v-row>

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
                <div v-if="errMsg.length">
                  <div v-for="err in errMsg" :key="err.idx">
                    <v-alert dismissible dense outlined type="error">
                      {{ err }}
                    </v-alert>
                  </div>
                </div>
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
          <v-btn
            v-if="isNewRegist"
            color="blue darken-1"
            text
            @click="createNewUser(input, mode)"
            >{{ mode }}</v-btn
          >
          <v-btn
            v-if="!isNewRegist"
            color="blue darken-1"
            text
            @click="login(input, mode)"
            >{{ mode }}</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- -->
  </v-layout>
</template>
<style scoped>
.mt-150 {
  margin-top: 150px;
}
.top {
  background-color: #f2f2f2;
}
</style>
<script>
const SUCCESSCODE = '00' // 処理成功コード
const FAILEDCODE = '80' // 処理失敗コード

const success = (d) => {
  return d.data.resultCode === SUCCESSCODE
}
const fail = (d) => {
  return d.data.resultCode === FAILEDCODE
}

const URLs = {
  login: 'http://localhost:8082/login',
  createNewUser: 'http://localhost:8082/createNewUser',
  trialLogin: 'http://localhost:8082/trialLogin',
}
// ユーザー入力初期状態
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
       * isLeapYear うるう年がどうかを判定します。
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

      // is30Days 月の日数が30日かどうかを判定します。
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
    // isNewRegist 新規登録であるかどうかを判定します。
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
      this.errMsg = []
    },
  },

  methods: {
    /**
     * createNewUser ユーザーの新規登録を行います。
     * 登録OK→次画面へ遷移。
     * 登録NG→エラーメッセージを表示。
     */
    createNewUser(input, mode) {
      // -- 実処理スタート --
      if (!this.validationOk(input, mode)) {
        // eslint-disable-next-line no-console
        console.log('input err')
        return
      }

      // 登録処理
      this.$axios
        .post(URLs.createNewUser, input)
        .then((res) => {
          console.log(res)
          if (success(res)) {
            console.log('success')
          }
          if (fail(res)) {
            this.errMsg.push(res.data.resultMessage)
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    /**
     * ログインボタン押下時に発火。
     * ログインOK→ユーザ情報を取得し、別画面へ遷移。
     * ログインNG→ログイン失敗の旨をユーザに伝える。
     */
    login(input, mode) {
      if (!this.validationOk(input, mode)) {
        // eslint-disable-next-line no-console
        console.log('input err')
        return
      }

      const params = {
        UserID: input.UserID,
        Password: input.Password,
      }

      // ログイン処理
      this.$axios
        .post(URLs.login, params)
        .then((res) => {
          console.log(res)
          if (success(res)) {
            console.log('success')
          }
          if (fail(res)) {
            console.log('failed')
            this.errMsg.push(res.data.resultMessage)
          }
        })
        .catch((err) => {
          console.error(err)
        })
    },
    /**
     * お試しログイン
     * お試しユーザの情報を取得し、次画面へ遷移します。
     */
    trialLogin() {
      // お試しログイン処理
      this.$axios
        .get(URLs.trialLogin)
        .then((res) => {
          console.log(res)
          if (success(res)) {
            console.log('success')
          }
          if (fail(res)) {
            console.log('failed')
            this.errMsg.push(res.data.resultMessage)
          }
        })
        .catch((err) => {
          console.error(err)
        })
      // TODO:次画面へ遷移
    },

    /**
     * モーダルの表示・非表示を制御します。
     */
    showModal(mode) {
      this.dialog = !this.dialog
      this.mode = mode
    },

    /**
     * validationOk
     * ユーザ入力値のバリデーションチェックを行います。
     * バリデーションエラー時はエラーメッセージを詰めます。
     */
    validationOk(input, mode) {
      // isAllInput 入力必須チェック
      const isAllInput = (input) => {
        let isValid = true
        Object.values(input).forEach((val) => {
          if (!val) {
            isValid = false
          }
        })
        if (!isValid) {
          this.errMsg.push('内容をすべて入力してください。')
        }
        return isValid
      }
      const isInputLoginInfo = (input) => {
        const isInputOK = input.UserID && input.Password
        if (!isInputOK) {
          this.errMsg.push('ID、Passwordを入力してください。')
        }
        return isInputOK
      }
      // validPassword パスワードチェック
      const validPassword = (input) => {
        // TODO: 3種類8文字以上のパスワードが入力されているかチェックする。
        const isInputOK =
          input.Password && input.Password === input.ConfirmPassword
        if (!isInputOK) {
          this.errMsg.push('パスワードは3種類8文字以上で入力してください。')
        }

        return isInputOK
      }

      // validEmail メールアドレスチェック
      const validEmail = (input) => {
        // TODO: 正規表現で、メールアドレスの形式かどうかチェックする。
        console.log(input)
        return true
      }

      if (mode === '新規登録') {
        // すべてのバリデーションに引っかからない場合であれば、trueを返却します。
        return isAllInput(input) && validPassword(input) && validEmail(input)
      }
      if (mode === 'ログイン') {
        return isInputLoginInfo(input)
      }
    },
  },
}
</script>
