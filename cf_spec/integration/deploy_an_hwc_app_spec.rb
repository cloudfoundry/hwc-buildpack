require 'open3'
require 'spec_helper'

describe 'CF HWC Buildpack' do
  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying an hwc app' do
    let(:app_name) { 'windows_app' }
    let(:app)      { Machete.deploy_app(app_name, stack: 'windows2012R2') }
    let(:browser)  { Machete::Browser.new(app) }

    context 'with a cached buildpack', :cached do
      it 'deploys without hitting the internet' do
        expect(app).to be_running

        # skip until moved into buildpacks-ci
        # expect(app).not_to have_internet_traffic
        expect(app).not_to have_logged(/Download \[/)
        expect(app).to have_logged(/Copy \[/)

        browser.visit_path('/')

        expect(browser.body).to include('hello i am nora')
      end
    end

    context 'with an uncached buildpack', :uncached do
      it 'deploys successfully' do
        expect(app).to be_running
        expect(app).to have_logged(/Download \[/)

        browser.visit_path('/')

        expect(browser.body).to include('hello i am nora')
      end
    end
  end
end
