using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Http;
using lab1.Repository.Interfaces;
using Entity;
using System;

namespace Controllers
{
    [ApiController]
    [Route("api/v1/persons")]
    public class PersonController: Controller
    {
        protected IPersonRepository personRepository;

        protected Person _person;
        public PersonController(IPersonRepository personRep)
        {
            personRepository = personRep;
        }

        /// <summary>
        /// Получить информацию о человеке по id
        /// </summary>
        /// <param name="Id">Идентификатор</param>
        /// <returns></returns>
        /// <response code="200" cref="Person">Человек</response>
        /// <response code="404">Не найдено соответствующего человека</response>
        /// <response code="500">Ошибка сервера</response>
        [IgnoreAntiforgeryToken]
        [HttpGet("{Id}")]
        public ActionResult GetPersonById(int Id)
        {
            try
            {
                Person person = personRepository.GetUserByID(Id);
                return person != null ? Ok(person) : NotFound();
            }
            catch (Exception e)
            {
                return StatusCode(500);
            }
        }

        /// <summary>
        /// Возвращает информацию о всех person
        /// </summary>
        /// <returns>Записи о всех person</returns>
        /// <response code="200" cref="Person">OK</response>
        /// <response code="404">Не найдено записей</response>
        /// <response code="500">Ошибка сервера</response>
        [IgnoreAntiforgeryToken]
        [HttpGet]
        public ActionResult GetAllPerson()
        {
            try
            {
                List<Person> persons = personRepository.GetAll();
                return persons != null && persons.Count() > 0 ? Ok(personRepository.GetAll()) : NotFound();
            }
            catch (Exception e)
            {
                return StatusCode(500);
            }
            }

        /// <summary>
        /// Создает запись о человеке
        /// </summary>
        /// <param name="person">Новая запись</param>
        /// <returns>Созданная новая запись о человеке</returns>
        /// <response code="201">Запись о человеке создана</response>
        /// <response code="404">Не удалось создать</response>
        /// <response code="500">Ошибка сервера</response>
        [IgnoreAntiforgeryToken]
        [HttpPost("")]
        //[ValidateAntiForgeryToken]
        public ActionResult PostPerson(Person person)
        {
            try
            {
                var person_ = personRepository.Add(person);
                return new CreatedResult("http://www.myapi.com/api/clients/" + person_.Id, null);
            }
            catch (Exception e)
            {
                return StatusCode(500);
            }

                }

        /// <summary>
        /// Изменяет запись о человеке на заданную
        /// </summary>
        /// <param name="person">Новая запись</param>
        /// <param name="Id">Идентификатор</param>
        /// <returns>При успешном выполнении возвращает новую запись</returns>
        /// <response code="200" cref="Person">Обновлено</response>
        /// <response code="404">Не найдено записи для обновления</response>
        /// <response code="500">Ошибка сервера</response>
        [IgnoreAntiforgeryToken]
        [HttpPatch("{Id}")]
        public ActionResult UpdatePersonById(int Id, [FromBody]Person person)
        {
            try
            {
                person.Id = Id;
                var person_ = personRepository.Update(person);
                return person_ != null ? Ok(person_) : NotFound();
            }
            catch (Exception e)
            {
                return StatusCode(500);
            }
        }

        /// <summary>
        /// Удаляет запись о человеке
        /// </summary>
        /// <param name="Id">Идентификатор</param>
        /// <returns>Код результата</returns>
        /// <response code="204" cref="Person">Успешно удалено</response>
        /// <response code="404">Не найдено записи</response>
        /// <response code="500">Ошибка сервера</response>
        [IgnoreAntiforgeryToken]
        [HttpDelete("{Id}")]
        public ActionResult DeletePersonById(long Id)
        {
            try
            {
                return personRepository.Delete(Id) > 0 ? NotFound() : Ok(null);
            }
            catch (Exception e)
            {
                return StatusCode(500);
            }
        }

    }
}
