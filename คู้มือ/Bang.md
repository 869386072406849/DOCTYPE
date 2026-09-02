<!-- wp:paragraph -->
<p>KBTG ใช้ Golang กับงานอะไร? มีอะไรควรรู้บ้างกับการใช้ Go ใน KBTG?</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p><a href="https://medium.com/@jirat_s?source=post_page---byline--dca091939147---------------------------------------"></a></p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p><a class="bx x by bz ca ab cb ac ae af ag ah ai ne" href="https://medium.com/@jirat_s?source=post_page---byline--dca091939147---------------------------------------">Jirat Srisawat</a>Follow</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>4 min read·Nov 10, 2022</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"z b cr u w"} -->
<p class="z b cr u w">Share</p>
<!-- /wp:paragraph -->

<!-- wp:image {"className":"qm qn qo qp qq qr qj qk paragraph-image"} -->
<figure class="wp-block-image qm qn qo qp qq qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*mMSqlXHLzxu1nNsma9fZmQ.jpeg" alt="Sticker แจก Golang guild members น้าา"/></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="327d">ในบทความนี้ ผมอยากจะมาอธิบายถึงงานที่ใช้ภาษา Go ใน KBTG โดยอยากเล่าตั้งแต่จุดเริ่มต้นของการนำภาษา Go มาใช้ที่ KBTG สู่ภาพการใช้งานในปัจจุบัน มีข้อดีข้อเสียอย่างไร และมีเรื่องอะไรที่ Developers ควรรู้บ้างในการพัฒนางานส่วนนั้นๆ</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"className":"rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv"} -->
<h2 class="wp-block-heading rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv" id="927f"><strong class="cb"><em class="so">จุดเริ่มต้นในการใช้ Go</em></strong></h2>
<!-- /wp:heading -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv" id="f22b">KBTG เริ่มนำ Go มาใช้เมื่อประมาณ 6 ปีก่อน (ช่วงปี 2016 ราวๆ Go 1.7 ซึ่งตอนนั้นยังไม่มี Go Module เลย ต้องใช้ Dep, Glide กันแบบลำบากพอสมควร) สาเหตุที่องค์กรใหญ่อย่างธนาคารกสิกรไทยและ KBTG กล้านำภาษาใหม่ๆ มาใช้ตอนนั้น เกิดจากการทำโปรเจค Innovation งานหนึ่งที่มีเงื่อนไขพิเศษหลายอย่าง ทั้ง Online และ Batch เช่น ต้องการลิมิต Memory แต่ละ Process ให้น้อยที่สุด ต้องการจัดการ Thread จำนวนมาก ต้องการ Speed ที่เร็ว ฯลฯ</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="2819">เพื่อตอบโจทย์โปรเจคดังกล่าว ทางทีมลอง Run Benchmark เทียบกันหลายๆ ภาษาตามที่คนในทีมคิดว่าเขียนและดูแลได้ ซึ่งตอนนั้นก็ได้หยิบ Go มาเปรียบเทียบโดยไม่ได้หวังอะไรมาก ปรากฏว่า Go สามารถทำได้ตามเงื่อนไขทั้งหมดและดันได้ผลดีกว่าภาษาอื่นๆ ด้วย แม้คนเขียนจะเพิ่งหัดเขียนและไม่เคยมีประสบการณ์มาก่อน ทำให้เราตัดสินใจเลือก Go สำหรับทำโปรเจคนั้นคู่กับ Stacks ใหม่ตัวอื่นๆ เช่น NoSQL, Microservice และ Containers พอเห็นว่าได้ผลดี โปรเจคอื่นๆ ก็เริ่มหันเข้าหา Golang และมีการใช้งานเพิ่มขึ้นเรื่อยๆ จนถึงทุกวันนี้</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"className":"rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv"} -->
<h2 class="wp-block-heading rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv" id="8223">Golang Trends in KBTG</h2>
<!-- /wp:heading -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*0n40fKNX8tYDQPFbmBpAPg.png" alt="จำนวน source code แต่ละภาษาย้อนหลัง 2 ปี"/><figcaption class="wp-element-caption">กราฟแสดงจำนวน Source Code แต่ละภาษาในแต่ละช่วงเวลา</figcaption></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="4abe">ข้อมูลเทรนด์นี้ดึงมาจาก Git ขององค์กร โดยเป็นการเปรียบเทียบจำนวน Source Code แต่ละภาษาย้อนหลังไป 2 ปี ตั้งแต่ปลายปี 2020 จนถึงกลางปี 2022 จะเห็นว่า Golang มีการใช้สูงขึ้นเรื่อยๆ จนเพิ่งขึ้นมาเป็นอันดับ 1 ได้ไม่นานมานี้เอง จากเดิมที่ระบบ Legacy ของธนาคารจะเป็น Java ซะส่วนใหญ่ ซึ่งเป็นหนึ่งในข้อพิสูจน์ว่าการใช้งานภาษา Go ใน KBTG นั้นเติบโตรวดเร็วมากภายในเวลาเพียง 4 ปี จากที่การใช้งานเป็นศูนย์ กลายเป็นภาษาที่มีใช้กันมากที่สุด</p>
<!-- /wp:paragraph -->

<!-- wp:quote {"className":"te tf tg"} -->
<blockquote class="wp-block-quote te tf tg"><!-- wp:paragraph {"className":"qx qy th qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="qx qy th qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="c6da">ปล. กราฟนี้ไม่ได้บอกว่าภาษาไหนดีกว่ากันนะ บอกแค่จำนวน Source Code ในองค์กรเฉยๆ โดยไม่ได้อ้างอิงถึงคุณภาพหรือความสามารถของแต่ละภาษา เพราะงั้นไม่น่ามีดราม่าเนอะ 555</p>
<!-- /wp:paragraph --></blockquote>
<!-- /wp:quote -->

<!-- wp:heading {"className":"rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv"} -->
<h2 class="wp-block-heading rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv" id="9dd4">แล้วใช้ Go กับงานอะไรบ้างล่ะ?</h2>
<!-- /wp:heading -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv" id="2691">อันนี้จะอธิบายยากหน่อย เพราะระบบในธนาคารค่อนข้างใหญ่ ซับซ้อน และแตกต่างกันมาก แต่เราก็พอจะแบ่งกลุ่มงาน Go ที่คล้ายๆ กันออกมาได้บ้าง โดยใช้วิธีแบ่งระบบของธนาคารเป็นส่วนย่อยๆ และเล่ารวมๆ ว่างาน Go แต่ละส่วนต่างกันยังไง</p>
<!-- /wp:paragraph -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*iLQnthY-Mi21qviaPx4XKg.jpeg" alt="high level architecture ส่วนที่เกี่ยวข้องกับ Go"/><figcaption class="wp-element-caption">จริงๆ ระบบข้างในซับซ้อนกว่านี้มาก แต่เพื่อให้ง่ายสำหรับบทความนี้ ขอเล่าแค่นี้พอ</figcaption></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="1c04">จากรูปข้างบน เราสามารถแบ่งระบบธนาคารเป็น 3 ส่วนใหญ่ๆ ดังนี้</p>
<!-- /wp:paragraph -->

<!-- wp:list {"className":""} -->
<ul class="wp-block-list"><!-- wp:list-item -->
<li><strong class="qz lo">Channel</strong> เช่น Mobile Banking, Web รวมถึง Back-office ที่พนักงานธนาคารใช้ หรือ API ที่ให้บริษัทลูกค้าหรือพาร์ทเนอร์ใช้</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li><strong class="qz lo">Backend &amp; Integration</strong> คือ ส่วน BFF (Backend for Frontend) ของแต่ละ Channel (ไม่ได้มี 1:1 ตามรูปนะ แต่ขึ้นอยู่กับดีไซน์ของแต่ละระบบ) รวมถึงระบบ Enterprise Integration Layer ที่ทำหน้าที่เฉพาะ อย่างเช่นการคุยกับ Core</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li><strong class="qz lo">Core</strong> เช่น ระบบ Core Banking ต่างๆ ที่เก็บบัญชีธนาคารหรือข้อมูลลูกค้าทั้งหมด</li>
<!-- /wp:list-item --></ul>
<!-- /wp:list -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="b98d">ระบบที่ใช้ Golang ส่วนใหญ่ใน KBTG จะอยู่ตรงช่วง Backend &amp; Integration ที่ต้องมีการรับ Request จาก Channel และส่งต่อไปยัง Core หรือดึงข้อมูลมา Process บางอย่างตรงงาน Batch ซึ่งงานแต่ละส่วนจะใช้ความสามารถของ Go แตกต่างกันไป ตามนี้</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3,"className":"tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv"} -->
<h3 class="wp-block-heading tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv" id="f057">Go as Backend</h3>
<!-- /wp:heading -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*uIppY6I0Z2GfCWn7fWiPTw.jpeg" alt=""/></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="55e8">ระบบที่ใช้ Go เป็น Backend มักจะใช้ Go ทำ RESTful API กับ Frontend ภาษาต่างๆ เช่น Swift, Kotlin, Flutter, JavaScript และทำ CRUD กับ Database ตามที่แต่ละระบบออกแบบไว้</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="149b">ความท้าทายของงานส่วนนี้ มีทั้งเรื่อง&nbsp;<strong class="qz lo">API Design, Scalability, Security&nbsp;</strong>และ&nbsp;<strong class="qz lo">Time to Market</strong>&nbsp;เนื่องจากเป็นระบบที่เน้นทำงานร่วมกับ Frontend ต่างๆ ที่มักมีการเปลี่ยนแปลง Requirements และ UI บ่อย จึงต้องดีไซน์ API ให้ยืดหยุ่นเพียงพอ ไปจนถึงการทำให้สเกลได้ในกรณีที่ระบบมีคนใช้งานสูงขึ้นแบบ Exponential ยังไม่รวม Security Requirements ต่างๆ ที่ต้องทำให้ปลอดภัยที่สุดตาม Standard ของธนาคารอีก</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="c010">ซึ่งจุดนี้เองที่ภาษา Go ช่วยให้การทำงานง่ายขึ้น ด้วยความ Simple ของภาษาและพวก Built-in Tools เช่น HTTP Server ในการ Implement API ได้ง่ายและเร็ว โดยไม่ต้องเขียนโค้ดหลายบรรทัด</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="322a">นี่จึงทำให้ Developers กลุ่มนี้มีความรู้ที่กว้างและหลากหลาย บางคนรู้ไปถึง Web หรือทำ Full Stack ได้ก็มี แถมส่วนใหญ่มักจะมีการใช้ Agile Methodology ในการพัฒนาระบบด้วย</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="8633">Technology หรือ Practice ที่ผู้พัฒนาระบบส่วนนี้ควรรู้ เช่น</p>
<!-- /wp:paragraph -->

<!-- wp:list {"className":""} -->
<ul class="wp-block-list"><!-- wp:list-item -->
<li>API Authorization &amp; Authentication</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Security, OWASP Secure Coding</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Database and Queries Optimization</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Microservices Design</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>OpenAPI spec, Swagger, gRPC</li>
<!-- /wp:list-item --></ul>
<!-- /wp:list -->

<!-- wp:heading {"level":3,"className":"tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv"} -->
<h3 class="wp-block-heading tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv" id="98d7">Go as Integration Layer</h3>
<!-- /wp:heading -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*G3PiLw36ZnfoAXrc9Ob8ng.jpeg" alt=""/></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="b784">Integration Layer ใน KBTG คือระบบกลางที่รวม Transaction จาก Backend แต่ละ Channel มาทำงานบางอย่าง เช่น การทำธุรกรรมทางการเงินทาง PromptPay เพื่อตัดเงินหรือเอาเงินเข้า เป็นต้น</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"className":"ui b uj uk ul um un"} -->
<h2 class="wp-block-heading ui b uj uk ul um un">Get&nbsp;Jirat Srisawat’s stories in&nbsp;your&nbsp;inbox</h2>
<!-- /wp:heading -->

<!-- wp:paragraph {"className":"z b cr u w"} -->
<p class="z b cr u w">Join Medium for free to get updates from&nbsp;this&nbsp;writer.Subscribe</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"z b cr u cv"} -->
<p class="z b cr u cv">Remember me for faster sign in</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rc rd re rg rh gx rj rk ha rm rn hd rp rq uc rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rc rd re rg rh gx rj rk ha rm rn hd rp rq uc rr kb cv" id="ac92">ความท้าทายของงานส่วนนี้จะเป็นเรื่อง&nbsp;<strong class="qz lo">Performance, Reliability</strong>&nbsp;และ&nbsp;<strong class="qz lo">Integration&nbsp;</strong>เพราะเป็นระบบกลางของทุกงาน ดังนั้นความเสถียรจึงเป็นเรื่องที่สำคัญที่สุด ถ้าระบบนี้ล่ม อาจถึงขั้นทำให้ทุกบริการของธนาคารใช้งานไม่ได้เลย นอกจากนี้ยังต้องออกแบบให้ Performance รองรับได้สูงที่สุดโดยไม่มี Error (นึกสภาพว่าถ้ามี Error ซัก 1 รายการ ก็เป็นปัญหาใหญ่แล้ว เช่น ลูกค้าโอนเงินแล้วเงินโดนตัด แต่ไม่เข้าบัญชีปลายทาง) บางกรณีอาจต้องมีการใช้ Event-driven ผ่าน MQ มาช่วยด้วย</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="fd71">ในมุมกลับกัน เรื่อง Integration กับภาษา Go ก็ถือเป็นงานที่ค่อนข้างยากกว่าภาษาอื่น เพราะแม้ Go จะเก่งเรื่อง Performance แต่ Go ไม่ Mature เลยในการ Integrate กับ Legacy Technology หรือ Proprietary ต่างๆ เช่น บาง Legacy ธนาคารยังต้องยิงด้วย SOAP หรือ Socket การต้องมาปั้น XML ยิง SOAP หรือ Socket ด้วยภาษา Go มีตัวช่วยน้อยกว่าภาษาอื่นมาก และยังไม่รวม Proprietary Database ที่เค้ามี ODBC/JDBC มาให้ ซึ่งใช้กับ Go ไม่ได้หรือไม่สมบูรณ์ ทำให้บางงานมีการออกแบบที่แปลกๆ เช่น เขียน Go ไปยิง Java ที่ต่อ JDBC หา Proprietary Database อีกที เรื่องนี้ทางเรามองเป็นข้อเสียของ Go เหมือนกัน</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="a373">ทั้งหมดนี้ทำให้ Developer ส่วน Integration จะมีความรู้และประสบการณ์ในเชิงลึกของเรื่องที่กล่าวมา และอาจจะต้องรู้ถึง Infrastructure ด้วย เพื่อให้จูน Performance ได้เต็มประสิทธิภาพ ดังนั้นสิ่งที่ควรรู้จะเป็นในเรื่องต่อไปนี้</p>
<!-- /wp:paragraph -->

<!-- wp:list {"className":""} -->
<ul class="wp-block-list"><!-- wp:list-item -->
<li>Infrastructure, Network, I/O</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>OS Parameters Optimization and Tuning</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Software Architecture Pattern e.g. SAGA, Event Driven Architecture, Service Mesh, Concurrency Pattern</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Monitoring &amp; Observability</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Exception Handling</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Profiling, Debugging</li>
<!-- /wp:list-item --></ul>
<!-- /wp:list -->

<!-- wp:heading {"level":3,"className":"tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv"} -->
<h3 class="wp-block-heading tr rt ln z ru gt ts ao gu gv tt aq gw gx tu gy gz ha tv hb hc hd tw he hf tx cv" id="b3d7">Go as Batch</h3>
<!-- /wp:heading -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*qpVISVFu0hVv54jIwXvgmg.jpeg" alt=""/></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="34ae">งาน Batch ใน KBTG มักจะเป็นการนำ Data มา Process โลจิคบางอย่าง เพื่อส่งต่อให้แก่ระบบ Downstream ต่อๆ ไป ซึ่งข้อจำกัดมักจะเป็นเรื่องของ&nbsp;<strong class="qz lo">ปริมาณ Data vs Runtime</strong></p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="5760">ความท้าทายของงานส่วนนี้คือเราจะเขียนโปรแกรมยังไงให้ทำงานได้เร็วที่สุดกับข้อมูลปริมาณมหาศาล? ซึ่งไม่พ้นการที่ต้องเข้าใจเรื่อง Algorithm และ Parallel Processing ดีพอที่จะ Optimize โปรแกรมได้ เช่น Batch บางตัวต้องนำข้อมูลธุรกรรมการเงินจากตลอดทั้งวัน (ประมาณ 20 ล้านรายการต่อวัน) มาคำนวณบางอย่างให้เสร็จภายใน 1 ชั่วโมง ก่อนส่งข้อมูลต่อไปยังระบบอื่น ถ้าทำงานช้า ก็อาจจะทำให้เปิดระบบธนาคารตอนเช้าไม่ทัน</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="263a">Developers กลุ่มนี้มักจะได้ใช้ Goroutine และ Channel ทำ Parallel Processing ซึ่งเป็นจุดขายของ Go อยู่แล้ว และอาจจะต้องออกแบบ Algorithm บางอย่างเองเพื่อ Optimize การทำงานให้เร็วที่สุด แถมยังต้องคำนึงถึง Exception และ Error Handling เป็นอย่างดี เพราะต้องสามารถ Recovery ได้ทันในเวลาที่กำหนด แม้จะเจอเคสแปลกๆ ที่คาดไม่ถึงเข้ามา</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="1bee">ตัวอย่างสิ่งที่ Developers กลุ่มนี้มักศึกษาและได้ใช้ในงาน เช่น</p>
<!-- /wp:paragraph -->

<!-- wp:list {"className":""} -->
<ul class="wp-block-list"><!-- wp:list-item -->
<li>Algorithm &amp; Data Structure</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Parallelism</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Text/DB Bulk Processing</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Error Handling &amp; Auto Recovery</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Reporting Tools</li>
<!-- /wp:list-item --></ul>
<!-- /wp:list -->

<!-- wp:heading {"className":"rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv"} -->
<h2 class="wp-block-heading rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv" id="7dc6">Common Practices</h2>
<!-- /wp:heading -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv" id="4942">แม้ว่า 3 กลุ่มตัวอย่างที่ยกมาข้างต้นจะแตกต่างกัน แต่จริงๆ ก็มี Technology &amp; Practice บางส่วนที่เป็น Standard ของ KBTG ที่ทุกทีมต้องประยุกต์ใช้อยู่ด้วย เช่น</p>
<!-- /wp:paragraph -->

<!-- wp:list {"className":""} -->
<ul class="wp-block-list"><!-- wp:list-item -->
<li>DevSecOps</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Unit Test, Code Review &amp; Code Style</li>
<!-- /wp:list-item -->

<!-- wp:list-item -->
<li>Observability &amp; Monitoring</li>
<!-- /wp:list-item --></ul>
<!-- /wp:list -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="ecae">ซึ่งภาษา Go สามารถช่วยให้เรื่องพวกนี้ง่ายขึ้น เพราะ Go เองก็มี Built-in Tools หลายตัว อย่าง go Test และ gofmt ที่จัดการบางเรื่องได้เลย ทำให้ไม่ค่อยมีปัญหาเรื่อง Compatibility เวลาใช้ร่วมกับซอฟต์แวร์อื่นๆ เช่น DevSecOps</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"className":"rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv"} -->
<h2 class="wp-block-heading rs rt ln z ru rv rw rx gu ry rz sa gw sb sc sd se sf sg sh si sj sk sl sm sn cv" id="01d7">สรุปงาน Golang กับ KBTG</h2>
<!-- /wp:heading -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra sp rc rd re sq rg rh gx sr rj rk ha ss rm rn hd st rp rq rr kb cv" id="e5b6">Golang เป็นภาษาที่ได้รับความนิยมมากขึ้นและมีการใช้ในหลายๆ ระบบของ KBTG เนื่องด้วยจุดเด่นหลายข้อ ไม่ว่าจะเป็นความ Simple ของภาษา, Built-in Tools, Performance และ Parallelism แต่ Golang ก็มีข้อเสียตรงที่ไม่เหมาะกับงานบางอย่าง เช่น งานที่เน้น Integration กับ Legacy หรือ Proprietary System ต่างๆ ของธนาคารเป็นจำนวนมาก</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="c5fb">บทความนี้ต้องการสื่อว่างานแต่ละงานมีการใช้จุดแข็งของ Go แตกต่างกัน และต้องอาศัยความรู้เสริมในเรื่องอื่นๆ ที่เกี่ยวข้องกับงานนั้นด้วย ซึ่ง Developers ที่สามารถใช้ความสามารถของภาษาคู่กับความรู้อื่นๆ ได้เหมาะสม จะทำให้สามารถเขียนโปรแกรมที่มีประสิทธิภาพดีที่สุดได้</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="246f">สุดท้ายนี้ผู้เขียนหวังว่าบทความนี้จะช่วยให้ Golang Developers เห็นภาพมากขึ้นว่า KBTG ใช้ Go กับงานส่วนไหนบ้าง และมีเรื่องอะไรที่ควรรู้บ้างในการพัฒนางานแต่ละส่วน ถ้าสงสัยหรือมีคำถามอะไร สามารถทิ้งคำถามไว้ได้เลย และจะพยายามตอบให้ครับ</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="75a8">ปล. ขอโปรโมทหน่อยว่า KBTG มี Golang Guild ที่เป็นเหมือนชมรมรวมคนที่สนใจ Go มาแบ่งปันความรู้ แลกเปลี่ยนประสบการณ์ และช่วยเหลือระหว่างทีมที่ใช้ Go ด้วยกัน ดังนั้นถ้าใครสนใจมาเป็น Golang Developer ใน KBTG ไม่เหงาแน่นอน เพราะเรามีคนที่เจออะไรคล้ายๆ กันพร้อมช่วยแก้ปัญหาอย่างเต็มที่</p>
<!-- /wp:paragraph -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*UOaT27UKAAyFQ4Ee2n4ZFg.jpeg" alt=""/></figure>
<!-- /wp:image -->

<!-- wp:image {"className":"sv sw sx sy sz qr qj qk paragraph-image"} -->
<figure class="wp-block-image sv sw sx sy sz qr qj qk paragraph-image"><img src="https://miro.medium.com/v2/resize:fit:1400/1*XljFE94yc2-CgJlhKCUTlQ.jpeg" alt=""/><figcaption class="wp-element-caption">ตัวอย่างที่ Golang Guild แชร์เรื่อง Observability ในงาน KBTG Inspire</figcaption></figure>
<!-- /wp:image -->

<!-- wp:paragraph {"className":"pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv"} -->
<p class="pw-post-body-paragraph qx qy ln qz b ra rb rc rd re rf rg rh gx ri rj rk ha rl rm rn hd ro rp rq rr kb cv" id="0c7d">สำหรับ Golang Developers คนไหนที่สนใจอยากมาร่วมทำงาน Go ด้วยกันที่ KBTG ตอนนี้เรายังเปิดรับอีกจำนวนมาก สามารถส่งใบสมัครมาได้ที่&nbsp;<a class="bx wv" href="https://www.kbtg.tech/" rel="noreferrer noopener" target="_blank">www.kbtg.tech</a></p>
<!-- /wp:paragraph -->
