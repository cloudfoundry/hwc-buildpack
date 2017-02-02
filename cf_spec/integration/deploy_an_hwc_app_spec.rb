require 'spec_helper'

describe 'CF HWC Buildpack' do
  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying an hwc app' do
    let(:app_name) { 'windows_app' }

    let(:app) { Machete.deploy_app(app_name, buildpack: 'hwc-test-buildpack') }
    let(:browser) { Machete::Browser.new(app) }

    it 'deploys successfully' do
      expect(app).to be_running

      browser.visit_path('/')

      expect(browser.body).to include('hello i am nora')
    end
  end
end
