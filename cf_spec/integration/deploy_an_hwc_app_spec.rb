require 'open3'
require 'spec_helper'

describe 'CF HWC Buildpack' do
  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying an hwc app' do
    let(:app_name) { 'windows_app' }
    let(:app)      { Machete.deploy_app(app_name, buildpack: 'hwc-test-buildpack') }
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

    context 'ginkgo specs' do
      before do
        @old_gopath = ENV['GOPATH']
        ENV['GOPATH'] = Dir.pwd
      end

      after do
        ENV['GOPATH'] = @old_gopath
      end

      it 'passes all tests' do
        _, stdout, stderr, wait_thr = Open3.popen3('ginkgo -r')
        exit_status = wait_thr.value
        unless exit_status.success?
          puts "gingko stdout:"
          puts stdout.read
          puts "ginkgo stderr:"
          puts stderr.read
        end
        expect(wait_thr.value).to eq(0)
      end
    end
  end
end
