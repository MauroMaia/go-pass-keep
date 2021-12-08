# go-pass-keeper

[![Create Release](https://github.com/MauroMaia/go-pass-keep/actions/workflows/go.yml/badge.svg)](https://github.com/MauroMaia/go-pass-keep/actions/workflows/go.yml)
[![CodeQL](https://github.com/MauroMaia/go-pass-keep/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/MauroMaia/go-pass-keep/actions/workflows/codeql-analysis.yml)

This project work as a case study to crete a cli password manager and increase 
my personal knowledge of golang.

## Install

+ Option 1: From RPM (WIP)
+ Option 2: From Github Releases
    ```bash
      cd /tmp &&
        wget https://github.com/MauroMaia/go-pass-keep/releases/download/v0.0.5/go-pass-keeper &&  
        chmod +x go-pass-keeper && 
        mv go-pass-keeper /usr/local/bin/
    ```  
+ Option 3: From git
    ```bash
        export GPK_DIR="/tmp/pgk" && (
          rm -rf $GPK_DIR
          git clone https://github.com/MauroMaia/go-pass-keep.git "$GPK_DIR"
          cd "$GPK_DIR"
          git checkout `git describe --abbrev=0 --tags --match "v[0-9]*" $(git rev-list --tags --max-count=1)`
        ) && cd "$GPK_DIR" && make
    ```

## Work in Progress

- [x] Save encrypted file
- [x] Load from encrypted file
- [ ] ... 
- [ ] Run a check if any of my password have been pawned 
- [ ] Random password generator 
- [ ] Table output format
- [ ] Export option (either way the main output  format is json)
- [ ] Create deployment process over rpm file
- [ ] Create a simple GUI version

## Licence

This project is developed under MIT license. You can consult the full licence text ![here](https://github.com/MauroMaia/go-pass-keep/blob/main/LICENSE).

Quick reminder:

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
