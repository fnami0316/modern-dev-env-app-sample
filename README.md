# 1. 概要
## 1.1. 前提
**Webアプリ開発におけるモダンな環境**  
昨今のWeb系の開発現場では、以下の要素を備えた環境を整えるのがモダンとされています。
- Gitでのコード管理(GitHub/GitLab etc.)
- クラウドインフラの利用(AWS/GCP/Azure etc.)
- クラウドインフラ設定をコード化(Terraform etc.)
- CI/CDパイプラインの構築(CircleCI/GitHub Actions etc.)
- Dockerの利用(ローカル開発環境の構築、CI/CD実行環境の構築、クラウドインフラ上でのサーバーのオートスケーリング対応 etc.)

システム構成を図にすると以下のような形になります。

![モダン開発環境-Lv2(抽象) drawio](https://github.com/user-attachments/assets/f046508e-ee23-4d89-942c-8b86f269a68a)

**モダンな環境を整えるメリット**
- アプリケーションの開発・テスト・リリースを複数人で円滑に遂行しやすくなる
- 優秀なエンジニアを採用しやすくなる
 
## 1.2. このリポジトリは何か？
モダンな環境を構築するためのサンプルプロジェクトです。  
このREADMEに含まれる[セットアップ手順](#2-セットアップの手順)の工程に従うことで、以下のシステム構成の環境を構築できます。  

### 1.2.1. システム構成

**システム構成図(全体)**  
![モダン開発環境-Lv2(具体)](https://github.com/user-attachments/assets/4574b6e8-2f31-40a3-b757-d819b44af822)

システムの構成要素|説明
---|---
Gitホスティングサービス|GitHub
クラウドインフラ|GCP:<br>- Cloud Run(コンテナサービス)<br>- Artifact Registry(Dockerレジストリ)<br>- Spanner(DB)<br>
Webアプリケーション|- 機能: gRPCのAPI。サンプルデータの追加、一覧取得、更新、削除を行うことができる。<br>- 実装方法:<br>  - 言語: Golang<br>  - フレームワーク: 未使用<br>  - 主なライブラリ: ProtocolBuffers, GORM
CI/CDツール|CircleCI
連絡ツール|Slack
ローカルマシン|Mac

**システム構成(ローカルマシン)**
![モダン開発環境-ローカル環境のシステム構成、コンテナオーケストレーションの設定](https://github.com/user-attachments/assets/ed475a50-47eb-4005-84e1-37f9341e72ea)


**システム構成(CircleCI, リモート環境)**
![モダン開発環境-Lv2(GCP, CircleCI) drawio](https://github.com/user-attachments/assets/20a1a292-cc8f-410e-bc3f-7fef7eb8a682)




### 1.2.2. 想定ワークフロー
上記のシステム構成において、以下に示すフローで開発・保守のタスクを進めることを想定しています。

※前提
- ブランチの種類
    - masterブランチ: リモート環境へデプロイする用意が整った状態のブランチ
    - 作業ブランチ: masterという名前でないブランチ。発生したタスクに対して変更を加えるためのブランチ
- Webアプリのデプロイとリリース: 互いに似ている言葉ですが、このサンプルプロジェクトにおいては、以下のように使い分けることにしています。
    - デプロイ: クラウド上でアプリケーションのコンテナを起動させること。非公開のURLを通してアクセス可能。
    - リリース: 公開しているURLへのアクセスへのトラフィックを切り替えて、通常利用できる状態にすること。

**ワークフロー図**  
![モダン開発環境-ワークフローLv2](https://github.com/user-attachments/assets/f9219bd5-290d-4ca3-b17f-43332d85e8f2)

### 1.2.3. 注意点
このサンプルプロジェクトでは、リモートに1つの環境があることしか想定していません。  
Webアプリケーション開発現場においては、QA環境、ステージング環境、本番環境といった環境も並行して存在しているのが一般的です。  
リモートに1つしか環境がない場合、例えばTerraformを適用するとリモート環境のクラウドインフラに反映されるため、インフラ周りの不具合が発生したら即流出という形になってしまいます。  
実際の現場でこのサンプルプロジェクトをベースに環境を構築する際には、対象のWebアプリのシステム検証〜本番リリースまでのワークフローをよく検討した上で、ブランチ運用、CircleCIの設定、Terraformの設定を変更することになると思います。  


# 2. セットアップの手順
※各サービス、ツールの経年変化により手順が多少変わる可能性があります。
 
## 2.1. 各サービス、ツールの初期設定
### 2.1.1. GitHub
#### GitHubのアカウントを作成
[GitHub](https://github.com/)へアクセス > Sign up
以降、画面の指示に従って作成してください。
 
### 2.1.2. GCP
#### Googleアカウントを作成
GCPインフラを構築するGoogleアカウントを用意してください。
 
#### GCPのFree Trialを申請
[GCPのFree Trial(90日間$300で使い放題)](https://cloud.google.com/free?hl=ja)を申請してください。
 
#### gcloud CLIをインストール
[gcloud CLIをインストールする](https://cloud.google.com/sdk/docs/install?hl=ja)に従ってインストールしてください。
 
### 2.1.3. Go
<TODO>
 
### 2.1.4. CircleCI
#### CircleCIのアカウントを作成
①[CircleCI](https://circleci.com/)へアクセス > Sign up > Sign up  
②任意のメールアドレス、パスワードを設定してアカウントを作成  
③Start a new organization - Get Startedを押下  
![circleci-start-new-org](https://github.com/user-attachments/assets/7a5a03b6-e752-4282-91f0-87ae3c9eea24)

④任意の組織名を入力してLet's goを押下  
 
### 2.1.5. Slack
#### Slackのアカウントを作成
[Slack](https://slack.com)へアクセス してアカウントを作成してください。
 
#### Slackのワークスペース、チャンネルを作成
GitHubやCircleCIからの通知先となるワークスペース、チャンネルを作成してください。
 
### 2.1.6. Docker
[https://docs.docker.com/desktop/install/mac-install/]からDocker for Desktopをインストールしてください。


### 2.1.7. Terraform
#### Terraformをインストール
```
# tfenvをHomebrewでインストール
brew install tfenv
  
# インストール可能なterraformのバージョンを確認
tfenv list-remote
  
# terraformをバージョンを指定してインストール
tfenv install <バージョン>
  
# インストール済みのterraformのバージョンを一覧表示(*付きが現在使用しているバージョン)
tfenv list
  
# 使用するterraformのバージョンを切り替え
tfenv use <バージョン>  
```
実務上、terreformのバージョンアップ作業が発生する可能性を考慮して、tfenv経由でインストールしています。
 
 
## 2.2. 各サービス、ツールの詳細設定
**前提** 
[2.1. 各サービス、ツールの初期設定](#21-各サービス、ツールの初期設定) の作業を終えていることが前提です。
 
### 2.2.1. GitHubリポジトリの作成と設定
#### 2.2.1.1. リポジトリの作成
GitHubへログイン > New repository > 以下の設定で Create repository
 
項目|値
---|---
Repository name|任意
Description|任意
Public/Private|Private
Add a README file|チェックしない
Add .gitignore|なし
Choose a license|None
 
 
#### 2.2.1.2. サンプルプロジェクトをリポジトリに登録
[サンプルプロジェクト](https://github.com/fnami0316/modern-dev-env-app-sample)をコピーしたリモートリポジトリをGitHubに作成します。
 
##### 手順
ローカルマシンのターミナルで以下を実行してください。
 
```
# サンプルプロジェクトをclone
git clone https://github.com/fnami0316/modern-dev-env-app-sample.git
 
# ローカルブランチ(master)のリモートリポジトリを、サンプルプロジェクトのものから自身のGitHubアカウントのリモートリポジトリへ変更
git remote set-url origin <自身のリモートリポジトリのURL>
 
# 現在のリモートリポジトリを確認
git remote -v
 
# 自身のGitHubアカウントのリモートリポジトリにサンプルプロジェクトをpush
git push origin
```
 
##### 参考
- [cloneしたリポジトリを別のリポジトリにpushする流れ](https://yuito-blog.com/repository-change/#index_id)
 
####  2.2.1.3. リポジトリの設定を変更
複数人で開発することを想定し、利便性向上、誤操作による復旧の手間を低減するという観点から、以下の項目を設定します。
- wiki利用可: デフォルト設定で可能
- masterブランチに対して以下の制限を設定(ブランチ保護ルール)
    - force pushを許可しない
    - 他のブランチをマージする前にPR必須にする
    - マージするブランチ側でのテスト必須にする
 
 
##### 手順
 
①GitHubへログイン > 対象のリポジトリのSettings > Code and automation - Rules で New ruleset  > New branch ruleset で New ruleset > New branch ruleset
→以下の画面が表示されます。
 
![image2024-7-5_16-49-30-](https://github.com/user-attachments/assets/f386aa20-b93f-4e6c-89d5-e3f811fb0a05)

 
②各項目を以下の通り設定してCreate押下
項目|　|値|補足
---|---|---|---
RulesetName| |任意のブランチ保護ルール名|ブランチ保護ルールを管理する上での名前。"masterブランチ"などわかりやすいものを設定。
Enforcement status| |Atcive|このブランチ保護ルールを適用するか否か
Targets|Target branches|master|ブランチ保護ルールを適用するブランチ名のパターン
Rules|Branch protections|- Require a pull request before merging: true<br>　- Required approvals: 1<br>　- Dismiss stale pull request approvals when new commits are pushed: true<br>　- Require review from Code Owners: false<br>　- Require approval of the most recent reviewable push: false<br>　- Require conversation resolution before merging: false<br>|- マージ前のPR必須にするか<br>　- 必須のapprove数<br>　- approve後にpushがあった場合、過去のapproveを却下する<br>　- コード所有者のapprove必須にするか<br>　- 直近のレビュー可能なpushのapproveを必須にするか<br>　 - マージ前の会話の解決を必須にするか<br>
　|　|- Require status checks to pass: true<br>　- Require branches to be up to date before merging: true<br>　- ci/circleci: build-and-test|- マージしようとしているブランチのステータスチェック通過を必須にするか<br>　- 最新のコードでチェックしなければならないか(後述のCircleCIとGitHubの連携設定が終わった後でないと設定不可)<br>　 
　|　|Block force pushes: true|
 
##### 参考 
- [プランと請求日を表示する](https://docs.github.com/ja/billing/managing-your-github-billing-settings/viewing-your-subscriptions-and-billing-date)
- [ウィキについて](https://docs.github.com/github/building-a-strong-community/about-wikis)
- [プルリクエストを自動的にマージする](https://docs.github.com/ja/pull-requests/collaborating-with-pull-requests/incorporating-changes-from-a-pull-request/automatically-merging-a-pull-request)
- [保護されたブランチについて](https://docs.github.com/ja/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches#restrict-who-can-push-to-matching-branches)
 
 
### 2.2.2. GitHub連携したCircleCIプロジェクトを作成
#### 手順
①CircleCIへログイン > Projects > CreateProject  
![circleci_create_project](https://github.com/user-attachments/assets/6d14082a-2159-44b9-9224-a0c50c02bec8)

 
②Build, test, and deploy a software application 押下  
![circleci_what_would_you_like](https://github.com/user-attachments/assets/1c325039-9971-40be-a32a-1dd1ff5d6b9f)

 
③GitHubを選択  
![circleci_select_repo](https://github.com/user-attachments/assets/99bd6f0f-df09-4c79-b1ed-12097bfc4982)


④連携するGitHubリポジトリのアカウントを選択 > Only select repositories にチェックを入れて、連携するリポジトリを選択して Install & Authorize  
![circleci_install_authorize](https://github.com/user-attachments/assets/d1e4c380-d102-4416-a7d1-746c4671d50e)

⑤Project Name へ任意のCircleCIプロジェクト名を入力

⑥Use Existing Configを押下  
![use_existing_config](https://github.com/user-attachments/assets/4cb83562-84b4-4a17-af29-cf0efb6e9d15)

⑦Start Building を押下  
![circleci-start-building](https://github.com/user-attachments/assets/05ba9085-2696-4d1a-87b5-1c40fc3ecb65)

この段階では、CircleCIとSlackを連携する設定をしていないため、Slack通知のjobでエラーになってしまいます。  
![circleci-error](https://github.com/user-attachments/assets/7a8b065a-95ee-4dd7-bddc-062275b885ea)

[2.2.3. CircleCIとSlackを連携](#223-CircleCIとSlackを連携)で設定します。
 
### 2.2.3. CircleCIとSlackを連携
#### 2.2.3.1. 通知先となるSlackチャンネルのワークスペースにSlackアプリを作成
##### 手順
**①アプリの作成** 
Slack API WebサイトのYour Apps > Create an App > From scratch > App Name に任意のアプリ名を入力して対象のワークスペースを選択　> Create App
 
**②アプリの権限設定** 
Settings-Basic Information > Add features and functionality > Permissions > Scopes - Bot Token Scopes > Add an OAuth Scope 押下、以下の選択を繰り返す
- chat:write
- chat:write.public
- files:write
 
**③アプリのインストール**
Features-OAuth & Permissions > Install to Workspace > 許可する > Slackアプリのインストール申請が承認された後、Features-OAuth & Permissions > Install to Workspace > 許可する
→ Bot User OAuth Token に トークンが表示される。次の工程で使うのでメモ。
 
 
#### 2.2.3.2. CircleCIの環境変数にSlackアプリのアクセストークンと通知先チャンネルを登録
##### 手順
CircleCIへログイン > [CircleCIのアカウントを作成](#CircleCIのアカウントを作成) で作成した組織を選択 > Projects > [2.2.2. GitHub連携したCircleCIプロジェクトを作成](#222-GitHub連携したCircleCIプロジェクトを作成)で作成したプロジェクトを選択 > Project Settings > Environment Variables より、以下2つの環境変数をそれぞれ登録
 
**Slackアプリのアクセストークン**
 
項目|値
---|---
Environment Variable Name|SLACK_ACCESS_TOKEN
Value|[2.2.3.1. 通知先となるSlackチャンネルのワークスペースにSlackアプリを作成](#2231-通知先となるSlackチャンネルのワークスペースにSlackアプリを作成)で作成したアクセストークン
 
**Slackチャンネル**
 
項目|値
---|---
Environment Variable Name|SLACK_DEFAULT_CHANNEL
Value|Slackのデスクトップアプリ上で、通知先のSlackチャンネルを右クリック > コピー > リンクをコピーで得られるURLの末尾の文字列(多分11文字)
 
#### 2.2.3.3. 動作確認
[2.2.2. GitHub連携したCircleCIプロジェクトを作成](#222-GitHub連携したCircleCIプロジェクトを作成)で失敗していたbuild-and-testワークフローを再実行して、Slackチャンネルに通知が届くことを確認する。  
![circleci-notification](https://github.com/user-attachments/assets/66240bf5-1478-4db1-9e2f-d6c8f206f9e7)

 
### 2.2.4. Terraformを使ってGCPのサービスをプロビジョニングする
#### 手順
ローカルマシンのターミナルで以下を実行してください。
 
```
# 1. terraform(Google Cloud SDK を使ったアプリケーション)のための認証
gcloud auth application-default login
  
# 2. ブラウザでGCPプロジェクトを操作するGoogleアカウントにログイン。ログイン後ターミナルに戻ってくる。
# ブラウザのクッキーを全削除しておかないとうまく動作しないかも
  
# 3. terraformのルートモジュールへ移動
cd deploy/terraform
  
# 4. terraform.backendをコメントアウト
  
# 5. terraformで必要になるproviderをセットアップ
terraform init
  
# 6. dry run
terraform plan
  
# 7. 適用(tfstateの保存先はローカル)
terraform apply
  
# 8. terraform.backendコメントアウトを元に戻す
  
# 9. dry run
terraform plan
  
# 10. 適用(tfstateの保存先がgcsに)
terraform apply
```

### 2.2.5. ローカル環境を起動する
#### 手順
ローカルマシンのターミナルで以下を実行してください。

```
# 1. プロジェクトルートへ移動
cd <プロジェクトルート>

# 2. ローカル環境を起動
make docker-compose-up-d

# 3. 動作確認(Goアプリケーションソースのテスト)
make docker-go-test

# 4. 動作確認(APIへリクエスト)
grpcurl -d '{"name": "sample1"}' localhost:8080 api.SampleService.CreateSample
```

# 3. ディレクトリ・ファイル構成
```
リポジトリルート/
    .circleci/
        config.yml # ../build/ci/config.ymlへのシンボリックリンクを設定。CircleCIの設定ファイルはここに配置されている必要がある。
    api/ # API仕様を配置するディレクトリ
        proto/ # 各RPCサービスの.protoを配置するディレクトリ
            sample.proto
    build/ # アプリのビルド、CIに必要なファイルを配置するディレクトリ
        ci/ # CIの設定やスクリプトを配置
            config.yml # CircleCIの設定ファイル
        packages/ # クラウド (AMI)、コンテナ (Docker)、OS (deb、rpm、pkg) パッケージの設定とスクリプトを配置
            docker/ # Dockerfileを配置
                Dockerfile.sample_app # サンプルプロジェクトのGoアプリケーションのイメージを生成するためのDockerfile
    cmd/ # アプリケーションのエントリポイントのソースコードを配置するディレクトリ
        sample_app/ # サンプルプロジェクトが提供するgRPC APIのエントリポイントのソースコード群を配置
            environment_variables.go # エントリポイントが必要とする環境変数の定義と、環境変数ロード処理を記述(エントリポイントから呼び出す)
            infrastructures.go # インフラ層の依存性注入処理を記述(エントリポイントから呼び出す)
            main.go # Goアプリケーションのエントリポイント
            presentations.go # プレゼンテーション層の依存性注入処理を記述(エントリポイントから呼び出す)
            usecases.go # アプリケーション層の依存性注入処理を記述(エントリポイントから呼び出す)
    deploy/ # アプリケーションをデプロイするインフラ周りの設定ファイルを配置
        docker-compose/ # docker-compose向けの設定ファイルを配置
            .env # docker-compose.yml上で展開する環境変数およびコンテナに渡す環境変数を記述
            docker-compose.yml # ローカル開発環境の構成管理
            docker-compose-circleci.yml # CircleCI環境の構成管理(ほぼローカル開発環境のものと同じだが、CircleCIの仕様上の制約で一部が異なる)
        terraform/ # terraformの設定ファイルを配置
            backend.tf   # tfstate保存先の設定
            main.tf      # エントリポイント。他の設定ファイルに記述していない全ての設定
            outputs.tf   # 全ての出力の宣言とその説明
            providers.tf # 使用する各providerのデフォルト設定
            README.md    # terraformのREADME
            variables.tf # 全ての変数の宣言とその説明
            versions.tf  # terraform自体のバージョンと各providerのバージョン設定
    internal/ # プライベートなアプリケーション, ライブラリのコードを配置するディレクトリ
        sample_app/
            application/ # アプリケーション層
                repository/ # リポジトリのインターフェース定義を配置(実装はインフラ層で)
                    transaction/ トランザクション関連のインターフェース定義を配置
                        connection_interface.go
                        transaction_interface.go
                    sample_repository_interface.go # Sampleエンティティリポジトリのインターフェース定義 
                    sample_repository_mock.go # Sampleエンティティリポジトリのインターフェースのモック(自動生成)
                request/ # 各ユースケースのリクエストパラメータの実装を配置(各パラメータは、ドメイン層のオブジェクトで扱う)
                    sample/ # Sampleサービスの各メソッドのリクエストパラメータの実装を配置
                        create_sample_request.go # SampleサービスCreateSampleメソッドのリクエストパラメータ
                        delete_sample_request.go # SampleサービスDeleteSampleメソッドのリクエストパラメータ
                        list_sample_request.go   # SampleサービスListSampleメソッドのリクエストパラメータ
                        update_sample_request.go # SampleサービスUpdateSampleメソッドのリクエストパラメータ                        
                response/ # 各ユースケースのレスポンスパラメータの実装を配置
                    sample/ # Sampleサービスの各メソッドのレスポンスパラメータの実装を配置(各パラメータは、ドメイン層のオブジェクトで扱う)
                        create_sample_response.go # SampleサービスCreateSampleメソッドのレスポンスパラメータ
                        delete_sample_response.go # SampleサービスDeleteSampleメソッドのレスポンスパラメータ
                        list_sample_response.go   # SampleサービスListSampleメソッドのレスポンスパラメータ
                        update_sample_response.go # SampleサービスUpdateSampleメソッドのレスポンスパラメータ               
                usecase/ # 各ユースケースの実行処理の実装を配置
                    sample/ # Sampleサービスの各メソッドの実行処理の実装を配置
                        create_sample_usecase.go           # ユースケース"SampleサービスCreateSampleメソッド"を実行する構造体の実装
                        create_sample_usecase_interface.go # ユースケース"SampleサービスCreateSampleメソッド"を実行する構造体のインターフェース定義
                        create_sample_usecase_mock.go      # ユースケース"SampleサービスCreateSampleメソッド"を実行する構造体のモック
                        delete_sample_usecase.go           # ユースケース"SampleサービスDeleteSampleメソッド"を実行する構造体の実装
                        delete_sample_usecase_interface.go # ユースケース"SampleサービスDeleteSampleメソッド"を実行する構造体のインターフェース定義
                        delete_sample_usecase_mock.go      # ユースケース"SampleサービスCreateSampleメソッド"を実行する構造体のモック
                        list_sample_usecase.go             # ユースケース"SampleサービスListSampleメソッド"を実行する構造体の実装
                        list_sample_usecase_interface.go   # ユースケース"SampleサービスListSampleメソッド"を実行する構造体のインターフェース定義
                        list_sample_usecase_mock.go        # ユースケース"SampleサービスListSampleメソッド"を実行する構造体のモック
                        update_sample_usecase.go           # ユースケース"SampleサービスUpdateSampleメソッド"を実行する構造体の実装
                        update_sample_usecase_interface.go # ユースケース"SampleサービスUpdateSampleメソッド"を実行する構造体のインターフェース定義
                        update_sample_usecase_mock.go      # ユースケース"SampleサービスUpdateSampleメソッド"を実行する構造体のモック
            domain/ # ドメイン層
                entity/ # エンティティの実装を配置
                    sample/ # sampleエンティティ関連の実装を配置
                        sample.go # sampleエンティティの実装
                service/ # ドメインサービスのソースコードを配置するディレクトリ
                value/ # 値オブジェクトのソースコードを配置するディレクトリ
                    sample_id.go # 値オブジェクトSampleIDのソースコード
                    sample_name.go # 値オブジェクトSampleNameのソースコード
            infrastructure/ # インフラ層
                repository/ # リポジトリの実装を配置
                    gorm/ # データストアがCloud Spanner、ORマッパー"GORM"利用したリポジトリの実装を配置(インターフェース定義はアプリケーション層でなされている)
                        transaction/ # トランザクション関連のインターフェース実装を配置
                            connection.go
                            transaction.go
                        sample_repository.go # Sampleエンティティのリポジトリインターフェースの実装を配置
                        setup.go # エントリポイントで呼び出すインフラ層の依存性注入処理の一部(GORMを利用してSpannerデータベースへ接続、スキーマ定義反映)
            presentation/ # プレゼンテーション層
                pb/ # protocで自動生成されるGo言語向けgRPCコードのファイル(*.pb.go) 置き場
                    api/                            
                        sample.pb.go # protocで自動生成
                        sample_grpc.pb.go # protocで自動生成
                sample/ # sample_grpc.pb.go で定義されているgRPCサーバーインターフェース"SampleSeviceServer"に対する実装を配置
                    create_sample.go # gRPCサーバーインターフェース"SampleServiceServer"のCreateSampleメソッドを扱う構造体の定義。
                    delete_sample.go # gRPCサーバーインターフェース"SampleServiceServer"のDeleteSampleメソッドを扱う構造体の定義。
                    list_sample.go   # gRPCサーバーインターフェース"SampleServiceServer"のListSampleメソッドを扱う構造体の定義。
                    sample.go        # gRPCサーバーインターフェース"SampleServiceServer"の実装にあたる構造体"SampleServiceServer"の定義。各メソッドは埋め込み構造体へ移譲
                    update_sample.go # gRPCサーバーインターフェース"SampleServiceServer"のUpdateSampleメソッドを扱う構造体の定義。
    scripts/ # Makefileから呼び出すスクリプトがあれば配置
    .editorconfig # 各ファイルの拡張子別のフォーマット設定(EditorConfigによる自動ファイルフォーマットは、多くのIDE・エディタで対応している)
    .gitignore
    go.mod
    go.sum
    makefile # よく使うがオプション指定が複雑なコマンドをエイリアスとして登録
    README.md # このサンプルプロジェクトの説明
```

# 4. タスク別の作業手順
[1.2.2. 想定ワークフロー](#122-想定ワークフロー)に登場する以下の工程について、具体的な作業手順を説明します。
- 作業ブランチに変更を加える
- クラウドインフラリソースの設定の変更を適用

## 4.1. Goアプリケーション(API)の開発・保守
<TODO>


## 4.2. クラウドインフラリソースの設定変更
1. リソースの変更内容に応じて、deploy/terraform/ 以下の.tfファイルを適宜修正します。
2. ローカルマシンのターミナルで以下を実行します。
 
```
# 1. Google Cloud SDKを使うアプリが、GCPプロジェクトのインフラリソースを変更する権限があるGoogleアカウントを使って認証できるようにする
gcloud auth application-default login

# ブラウザで対象のGoogleアカウントにログイン。ログイン後ターミナルに戻ってくる(ブラウザのクッキーを全削除しておかないとうまく動作しないかも)
  
# 2. terraformのルートモジュールへ移動
cd <プロジェクトルート>/deploy/terraform

# 3. .tfの構文に問題がないかチェック
terraform validate
  
# 4. 必要なproviderをセットアップ
terraform init
  
# 5. リソース設定の変更箇所を事前確認
terraform plan
  
# 6. リソース設定の変更を適用
terraform apply
```

## 4.3. CircleCIの設定変更
1. .circleci/config.yml を開き、設定を変更して作業ブランチにcommit、pushします。
2. [2.2.2. GitHub連携したCircleCIプロジェクトを作成](#222-github連携したcircleciプロジェクトを作成)で作成していたCircleCIのプロジェクトへアクセスし、期待通りの挙動になっていることを確認します。



